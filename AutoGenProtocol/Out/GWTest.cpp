#include "GWTest.h"

BOOL GWTest::Read( SocketInputStream& iStream )
{
__ENTER_FUNCTION
    iStream.Read(Value1, sizeof(INT));
    iStream.GetArray(Value2, Value2Count);
    iStream.Read(Value2Count, sizeof(SHORT));
    iStream.Read(Value3, sizeof(SHORT));
    iStream.Read(Value4, sizeof(BYTE));
    iStream.GetArray(Value5, Value5Count);
    iStream.Read(Value5Count, sizeof(SHORT));
    return TRUE;
__LEAVE_FUNCTION
    return FALSE;
}

BOOL GWTest::Write( SocketOutputStream& oStream ) const
{
__ENTER_FUNCTION
    oStream.Write(Value1, sizeof(INT));
    oStream.PutArray(Value2, Value2Count);
    oStream.Write(Value2Count, sizeof(SHORT));
    oStream.Write(Value3, sizeof(SHORT));
    oStream.Write(Value4, sizeof(BYTE));
    oStream.PutArray(Value5, Value5Count);
    oStream.Write(Value5Count, sizeof(SHORT));
    return TRUE;
__LEAVE_FUNCTION
    return FALSE;
}

UINT GWTest::Execute(Player* pPlayer)
{
__ENTER_FUNCTION   
    return GWTestHandler::Execute(this, pPlayer);
__LEAVE_FUNCTION
    return FALSE;
}