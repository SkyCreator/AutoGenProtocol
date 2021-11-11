using UnityEngine;
namespace ProjectS
{
    public class GCTest : PADownwardPacketBase
    {
        public override int Id
        {
            get
            {
                return 2883;
            }
        }
        
        public int m_Score;
        public short m_FlagCount;
        public ulong[] m_Flag = new ulong[5];
        public short m_Count;
        public short m_Value4Count;
        public byte[] m_Value4 = new byte[7];
        public WorldPostion m_wpPos;
        public Vector3 m_Dir;
        public short m_PosCount;
        public Vector2[] m_Pos = new Vector2[6];
        public short m_TitleCount;
        public string m_Title;
        public override void Clear()
        {
            m_Score = default(int);
            m_FlagCount = default(short);
            for ( int i=0; i<5; ++i )
            {
                m_Flag[i] = default(ulong);
            }
            m_Count = default(short);
            m_Value4Count = default(short);
            for ( int i=0; i<7; ++i )
            {
                m_Value4[i] = default(byte);
            }
            m_wpPos.Clear();
            m_Dir = default(Vector3);
            m_PosCount = default(short);
            for ( int i=0; i<6; ++i )
            {
                m_Pos[i] = default(Vector2);
            }
            m_TitleCount = default(short);
            m_Title = string.Empty;
        }
        public override int ReadStream(byte[] stream, int index, int size)
        {
            index = ReadInt32(out m_Score, stream, index, size);
            index = ReadInt16(out m_FlagCount, stream, index, size);
            for ( int i=0; i<m_FlagCount; ++i )
            {
                index = ReadUInt64(out m_Flag[i], stream, index, size);
            }
            index = ReadInt16(out m_Count, stream, index, size);
            index = ReadInt16(out m_Value4Count, stream, index, size);
            index = ReadBytes(out m_Value4, m_Value4Count, stream, index, size);
            index = ReadSingle(out m_wpPos.x, stream, index, size);
            index = ReadSingle(out m_wpPos.z, stream, index, size);
            index = ReadSingle(out m_Dir.x, stream, index, size);
            index = ReadSingle(out m_Dir.y, stream, index, size);
            index = ReadSingle(out m_Dir.z, stream, index, size);
            index = ReadInt16(out m_PosCount, stream, index, size);
            for ( int i=0; i<m_PosCount; ++i )
            {
                index = ReadSingle(out m_Pos[i].x, stream, index, size);
                index = ReadSingle(out m_Pos[i].y, stream, index, size);
            }
            index = ReadInt16(out m_TitleCount, stream, index, size);
            int len = PASocketStreamAssist.GetStringBytesLen(m_Title);
            index = ReadString(out m_Title, len, stream, index);
            return index;
        }
    }
}