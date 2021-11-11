#include "GWTest.h"
#include "Log.h"
#include "ServerPlayer.h"
#include "ServerManager.h"

UINT GWTestHandler::Execute(GWTest* pPacket, Player* pPlayer)
{
__ENTER_FUNCTION
    Assert(pPacket);
    ServerPlayer* pServerPlayer = (ServerPlayer)pPlayer;
    Assert(pServerPlayer);
    //写下你的代码
    return PACKET_EXE_CONTINUE;
__LEAVE_FUNCTION
    return PACKET_EXE_ERROR;
}
    