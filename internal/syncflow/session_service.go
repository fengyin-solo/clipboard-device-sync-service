package syncflow

import "clipboard/internal/session"

type SessionService struct {
	pool *session.Pool
}

func NewSessionService(pool *session.Pool) *SessionService {
	return &SessionService{pool: pool}
}

// Snapshot 返回一个闭包，供调用方在稍后的异步路径中读取当前设备会话。
//
// 复用污染的根因在于 sync.Pool 的复用语义：若在同步阶段就把 *Value 归还到池中，
// 而返回的闭包仍持有该指针，那么在异步读取真正发生之前，另一台设备的
// Acquire 可能复用同一个对象，把 DeviceID/Payload 改成新值；空入参时
// 还会残留上一次的 Payload。于是异步读到的设备与内容都变成了 B。
//
// 修复：在 Release 之前先做深拷贝生成不可变快照，闭包只持有这份副本，
// 与池对象彻底解耦。异步读取拿到的永远是调用时刻的数据，不受后续复用影响。
func (s *SessionService) Snapshot(deviceID string, payload []byte) func() session.Value {
	value := s.pool.Acquire(deviceID, payload)
	// 必须在 Release 之前完成深拷贝，使快照与池对象彻底解耦：
	//   - DeviceID 是 string，值拷贝即不可变；
	//   - Payload 是切片，仅做结构体值拷贝会共享底层数组；归还后下一个
	//     Acquire 会复用同一底层数组并覆盖内容，异步读取就会读到被污染的数据。
	//   因此 Payload 必须复制底层数组。
	snapshot := session.Value{
		DeviceID: value.DeviceID,
		Payload:  append([]byte(nil), value.Payload...),
	}
	s.pool.Release(value)
	return func() session.Value { return snapshot }
}
