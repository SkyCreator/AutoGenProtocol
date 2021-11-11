#ifndef __GWTest_H__
#define __GWTest_H__

#include "Type.h"
#include "PacketFactoryBaseWithDestoryPacket.h"
namespace Packets
{
    class GWTest : public Packet
    {
    public:
        GWTest(){CleanUp();}
        virtual ~GWTest(){};
        VOID CleanUp(){};

    //公共继承接口
    virtual BOOL Read(SocketInputStream& iStream);
    virtual BOOL Write(SocketOutputStream& oStream) const;
    virtual UINT Execute(Player* pPlayer);
    virtual PacketID_t GetPacketID() const { return PACKET_GW_TEST; }
    virtual UINT GetPacketSize() const 
    { 
        UINT totalSize = 0;
        totalSize += sizeof(INT);
        totalSize += sizeof(SHORT);
        totalSize += sizeof(UINT) * Value2Count;
        totalSize += sizeof(SHORT);
        totalSize += sizeof(BYTE);
        totalSize += sizeof(SHORT);
        totalSize += sizeof(CHAR) * Value5Count;
        return totalSize;
    }
    
    public:
            VOID SetValue1( INT v)
            {
                Value1 = v;
            }
            INT GetValue1(VOID) const 
            { 
                return Value1; 
            }
            SHORT* GetValue2Count()
            {
                return &Value2Count;
            } 
            VOID SetValue2( int i, UINT v)
            {
                Value2[i] = v;
            }
            UINT GetValue2(int i) const 
            { 
                return Value2[i]; 
            }
            VOID SetValue3( SHORT v)
            {
                Value3 = v;
            }
            SHORT GetValue3(VOID) const 
            { 
                return Value3; 
            }
            VOID SetValue4( BYTE v)
            {
                Value4 = v;
            }
            BYTE GetValue4(VOID) const 
            { 
                return Value4; 
            }
            SHORT* GetValue5Count()
            {
                return &Value5Count;
            } 
            VOID SetValue5( const CHAR* v)
            {
                if (NULL == v) {return;}
                INT len = sizeof(Value5) - 1;
                Value5[len] = '\0';
                strncpy( Value5, v, len );
                Value5Count = strlen(Value5);
            }
            const CHAR* GetValue5(VOID) 
            { 
                return Value5; 
            } 
    private:
            INT Value1;
            SHORT Value2Count;
            UINT Value2[5];
            SHORT Value3;
            BYTE Value4;
            SHORT Value5Count;
            CHAR Value5[10];
    };

    class GWTestFactory : public PacketFactoryBaseWithDestoryPacket
    {
    public:
        Packet* CreatePacket() { return new GWTest();}
        PacketID_t GetPacketID() const { return PACKET_GW_TEST; }
        UINT GetPacketMaxSize() const 
        {
            UINT totalSize = 0;
            totalSize += sizeof(INT);
            totalSize += sizeof(SHORT);
            totalSize += sizeof(UINT)* 5;
            totalSize += sizeof(SHORT);
            totalSize += sizeof(BYTE);
            totalSize += sizeof(SHORT);
            totalSize += sizeof(CHAR)* 10;
            return totalSize;
        }
    };

    class GWTestHandler
    {
    public:
        static UINT Execute( GWTest* pPacket, Player* pPlayer );
    };
}

using namespace Packets;

#endif
