#ifndef __CGTest_H__
#define __CGTest_H__

#include "Type.h"
#include "PacketFactoryBaseWithDestoryPacket.h"
namespace Packets
{
    class CGTest : public Packet
    {
    public:
        CGTest(){CleanUp();}
        virtual ~CGTest(){};
        VOID CleanUp(){};

    //公共继承接口
    virtual BOOL Read(SocketInputStream& iStream);
    virtual BOOL Write(SocketOutputStream& oStream) const;
    virtual UINT Execute(Player* pPlayer);
    virtual PacketID_t GetPacketID() const { return PACKET_CG_TEST; }
    virtual UINT GetPacketSize() const 
    { 
        UINT totalSize = 0;
        totalSize += sizeof(INT);
        totalSize += sizeof(UINT) * FlagCount;
        totalSize += sizeof(SHORT);
        totalSize += sizeof(SHORT);
        totalSize += sizeof(BYTE);
        totalSize += sizeof(CHAR) * NameCount;
        totalSize += sizeof(SHORT);
        totalSize += sizeof(GUID64_t);
        return totalSize;
    }
    
    public:
            VOID SetScore( INT v)
            {
                Score = v;
            }
            INT GetScore(VOID) const 
            { 
                return Score; 
            } 
            VOID SetFlag( int i, UINT v)
            {
                Flag[i] = v;
            }
            UINT GetFlag(int i) const 
            { 
                return Flag[i]; 
            }
            SHORT* GetFlagCount()
            {
                return &FlagCount;
            }
            VOID SetCount( SHORT v)
            {
                Count = v;
            }
            SHORT GetCount(VOID) const 
            { 
                return Count; 
            }
            VOID SetValue4( BYTE v)
            {
                Value4 = v;
            }
            BYTE GetValue4(VOID) const 
            { 
                return Value4; 
            } 
            VOID SetName( const CHAR* v)
            {
                if (NULL == v) {return;}
                INT len = sizeof(Name) - 1;
                Name[len] = '\0';
                strncpy( Name, v, len );
                NameCount = strlen(Name);
            }
            const CHAR* GetName(VOID) 
            { 
                return Name; 
            } 
            SHORT* GetNameCount()
            {
                return &NameCount;
            }
            GUID64_t* GetGuildID()
            {
                return &GuildID;
            }
    private:
            INT Score;
            UINT Flag[5];
            SHORT FlagCount;
            SHORT Count;
            BYTE Value4;
            CHAR Name[10];
            SHORT NameCount;
            GUID64_t GuildID;
    };

    class CGTestFactory : public PacketFactoryBaseWithDestoryPacket
    {
    public:
        Packet* CreatePacket() { return new CGTest();}
        PacketID_t GetPacketID() const { return PACKET_CG_TEST; }
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
            totalSize += sizeof(GUID64_t);
            return totalSize;
        }
    };

    class CGTestHandler
    {
    public:
        static UINT Execute( CGTest* pPacket, Player* pPlayer );
    };
}

using namespace Packets;

#endif
