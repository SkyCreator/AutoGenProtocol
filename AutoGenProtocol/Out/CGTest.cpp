#include "CGTest.h"

BOOL CGTest::Read( SocketInputStream& iStream )
{
__ENTER_FUNCTION
    iStream.Read(Score, sizeof(INT));
    iStream.GetArray(Flag, FlagCount);
    iStream.Read(FlagCount, sizeof(SHORT));
    iStream.Read(Count, sizeof(SHORT));
    iStream.Read(Value4, sizeof(BYTE));
    iStream.GetArray(Name, NameCount);
    iStream.Read(NameCount, sizeof(SHORT));
    GuildID.Read(iStream);
    return TRUE;
__LEAVE_FUNCTION
    return FALSE;
}

BOOL CGTest::Write( SocketOutputStream& oStream ) const
{
__ENTER_FUNCTION
    oStream.Write(Score, sizeof(INT));
    oStream.PutArray(Flag, FlagCount);
    oStream.Write(FlagCount, sizeof(SHORT));
    oStream.Write(Count, sizeof(SHORT));
    oStream.Write(Value4, sizeof(BYTE));
    oStream.PutArray(Name, NameCount);
    oStream.Write(NameCount, sizeof(SHORT));
    GuildID.Write(oStream);
    return TRUE;
__LEAVE_FUNCTION
    return FALSE;
}

UINT CGTest::Execute(Player* pPlayer)
{
__ENTER_FUNCTION   
    return CGTestHandler::Execute(this, pPlayer);
__LEAVE_FUNCTION
    return FALSE;
}