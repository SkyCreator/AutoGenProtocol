#include "WGTest.h"

BOOL WGTest::Read( SocketInputStream& iStream )
{
__ENTER_FUNCTION
    iStream.Read(Value1, sizeof(INT));
    iStream.GetArray(Value2, Value2Count);
    iStream.Read(Value2Count, sizeof(SHORT));
    iStream.Read(Value3, sizeof(SHORT));
    iStream.Read(Value4, sizeof(BYTE));
    iStream.GetArray(Value5, Value5Count);
    iStream.Read(Value5Count, sizeof(SHORT));
    iStream.Read(UserData, sizeof(FULLUSERDATA));
    return TRUE;
__LEAVE_FUNCTION
    return FALSE;
}

BOOL WGTest::Write( SocketOutputStream& oStream ) const
{
__ENTER_FUNCTION
    oStream.Write(Value1, sizeof(INT));
    oStream.PutArray(Value2, Value2Count);
    oStream.Write(Value2Count, sizeof(SHORT));
    oStream.Write(Value3, sizeof(SHORT));
    oStream.Write(Value4, sizeof(BYTE));
    oStream.PutArray(Value5, Value5Count);
    oStream.Write(Value5Count, sizeof(SHORT));
    oStream.Write(UserData, sizeof(FULLUSERDATA));
    return TRUE;
__LEAVE_FUNCTION
    return FALSE;
}

UINT WGTest::Execute(Player* pPlayer)
{
__ENTER_FUNCTION   
    return WGTestHandler::Execute(this, pPlayer);
__LEAVE_FUNCTION
    return FALSE;
}