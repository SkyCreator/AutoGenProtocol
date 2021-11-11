using UnityEngine;
namespace ProjectS
{
    public class CGTest : GameServerUpwardPacket
    {
        public override int Id
        {
            get
            {
                return 2882;
            }
        }
        
        public int m_Score;
        public short m_FlagCount;
        public ulong[] m_Flag = new ulong[5];
        public short m_Count;
        public short m_NameCount;
        public string m_Name;
        public short m_Value4Count;
        public byte[] m_Value4 = new byte[7];
        public WorldPostion m_wpPos;
        public Vector3 m_Dir;
        public Vector2 m_Pos;
        public override void Clear()
        {
            m_Score = default(int);
            m_FlagCount = default(short);
            for ( int i=0; i<5; ++i )
            {
                m_Flag[i] = default(ulong);
            }
            m_Count = default(short);
            m_NameCount = default(short);
            m_Name = string.Empty;
            m_Value4Count = default(short);
            for ( int i=0; i<7; ++i )
            {
                m_Value4[i] = default(byte);
            }
            m_wpPos.Clear();
            m_Dir = default(Vector3);
            m_Pos = default(Vector2);
        }
        public override int WriteStream(byte[] stream, int index, int size)
        {
            index = WriteInt32(m_Score, stream, index, size);
            index = WriteInt16(m_FlagCount, stream, index, size);
            for ( int i=0; i<m_FlagCount; ++i )
            {
                index = WriteUInt64(m_Flag[i], stream, index, size);
            }
            index = WriteInt16(m_Count, stream, index, size);
            index = WriteInt16(m_NameCount, stream, index, size);
            int len = PASocketStreamAssist.GetStringBytesLen(m_Name);
            index = WriteByte((byte)len, stream, index, size);
            index = WriteCharArray(m_Name.ToCharArray(), len, stream, index, size, true, true);
            index = WriteInt16(m_Value4Count, stream, index, size);
            index = WriteBytes(m_Value4, m_Value4Count, stream, index, size);
            index = WriteSingle(m_wpPos.x, stream, index, size);
            index = WriteSingle(m_wpPos.z, stream, index, size);
            index = WriteSingle(m_Dir.x, stream, index, size);
            index = WriteSingle(m_Dir.y, stream, index, size);
            index = WriteSingle(m_Dir.z, stream, index, size);
            index = WriteSingle(m_Pos.x, stream, index, size);
            index = WriteSingle(m_Pos.y, stream, index, size);
            return index;
        }
    }
}