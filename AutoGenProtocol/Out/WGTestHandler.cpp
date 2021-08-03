#include "WGTest.h"
#include "Log.h"
#include "GamePlayer.h"
#include "Scene.h"
#include "Obj_Human.h"
#include "ServerManager.h"

UINT WGTestHandler::Execute( WGTest* pPacket, Player* pPlayer )
{
__ENTER_FUNCTION
    if ( pPacket == NULL || pPlayer == null )
    {
        return PACKET_EXE_CONTINUE;
    }
    if ( pPacket->GetFromWhere() != PACKET_FROM_WORLD )
    {
        return PACKET_EXE_CONTINUE;
    }
    PlayerID_t PlayerID = pPacket->GetPlayerID();
    GamePlayer* pGamePlayer = g_pPlayerPool->GetPlayer(PlayerID);
    if ( pGamePlayer == NULL )
    {
        Assert(FALSE);
        return PACKET_EXE_CONTINUE;
    }
    Obj_Human* pHuman = pGamePlayer->GetHuman();
    Assert(pHuman);
    if ( !pGamePlayer->IsCanLogic() )
    {
        CacheGuid64Log(LOG_FILE_1, "ERROR WGTestHandler::pGamePlayer->IsCanLogic():GUID=%s Name:%s.", Guid2StringObj(pGamePlayer->m_HumanGUID).GetString(), pHuman->GetName());
        return PACKET_EXE_CONTINUE;
    }
    Scene* pScene = pHuman->getScene();
    if ( pScene == NULL )
    {
        Assert(FALSE);
        return PACKET_EXE_ERROR;
    }
    if ( ((Player_AtServer*)pPlayer)->IsServerPlayer() )
    {
        Assert(g_pServerManager->VerifyExecuteThread());

        CacheGuid64Log(LOG_FILE_1, "WGTestHandler::Execute(). GUID=%s.", Guid2StringObj(pHuman->GetGUID()).GetString());
        pScene->PushAsyncPacket( pPacket, pPacket->GetPlayerID() );
        return PACKET_EXE_CONTINUE;
    }
    else if ( ((Player_AtServer*)pPlayer)->IsGamePlayer() )
    {
        if ( !pHuman->IsCanLogic() )
        {
            CacheGuid64Log(LOG_FILE_1, "ERROR WGTestHandler::IsCanLogic():GUID=%s Name:%s.", Guid2StringObj(pGamePlayer->m_HumanGUID).GetString(), pHuman->GetName());
            return PACKET_EXE_CONTINUE;
        }
    }
    
    //写下你的代码
    
    return PACKET_EXE_CONTINUE;
__LEAVE_FUNCTION
    return PACKET_EXE_ERROR;
}