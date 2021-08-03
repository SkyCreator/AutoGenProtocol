#ifndef __GCTest_H__
#define __GCTest_H__

#include "Type.h"
#include "PacketFactoryBaseWithDestoryPacket.h"
namespace Packets
{
    class GCTest : public Packet
    {
    public:
        GCTest(){CleanUp();}
        virtual ~GCTest(){};
        VOID CleanUp(){};

    //公共继承接口
    virtual BOOL Read(SocketInputStream& iStream);
    virtual BOOL Write(SocketOutputStream& oStream) const;
    virtual UINT Execute(Player* pPlayer);
    virtual PacketID_t GetPacketID() const { return PACKET_GC_TEST; }
    virtual UINT GetPacketSize() const 
    { 
        UINT totalSize = 0;
        totalSize += sizeof(INT);
        totalSize += sizeof(UINT) * Value2Count;
        totalSize += sizeof(SHORT);
        totalSize += sizeof(SHORT);
        totalSize += sizeof(BYTE);
        totalSize += sizeof(CHAR) * Value5Count;
        totalSize += sizeof(SHORT);
        totalSize += sizeof(WORLD_POS_3D);
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
            VOID SetValue2( int i, UINT v)
            {
                Value2[i] = v;
            }
            UINT GetValue2(int i) const 
            { 
                return Value2[i]; 
            }
            SHORT* GetValue2Count()
            {
                return &Value2Count;
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
            SHORT* GetValue5Count()
            {
                return &Value5Count;
            }
            WORLD_POS_3D* GetPosWorld()
            {
                return &PosWorld;
            }
    private:
            INT Value1;
            UINT Value2[5];
            SHORT Value2Count;
            SHORT Value3;
            BYTE Value4;
            CHAR Value5[10];
            SHORT Value5Count;
            WORLD_POS_3D PosWorld;
    };

    class GCTestFactory : public PacketFactoryBaseWithDestoryPacket
    {
    public:
        Packet* CreatePacket() { return new GCTest();}
        PacketID_t GetPacketID() const { return PACKET_GC_TEST; }
        UINT GetPacketMaxSize() const 
        {
            UINT totalSize = 0;
            totalSize += sizeof(INT);
            totalSize += sizeof(UINT)* 5;
            totalSize += sizeof(SHORT);
            totalSize += sizeof(SHORT);
            totalSize += sizeof(BYTE);
            totalSize += sizeof(CHAR)* 10;
            totalSize += sizeof(SHORT);
            totalSize += sizeof(WORLD_POS_3D);
            return totalSize;
        }
    };

    class GCTestHandler
    {
    public:
        static UINT Execute( GCTest* pPacket, Player* pPlayer );
    };
}

using namespace Packets;

#endif
