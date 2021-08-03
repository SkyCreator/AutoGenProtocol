#include "CGTest.h"
#include "Log.h"
#include "GamePlayer.h"
#include "Scene.h"
#include "Obj_Human.h"

UINT CGTestHandler::Execute( CGTest* pPacket, Player* pPlayer )
{
__ENTER_FUNCTION
    Assert(pPacket != NULL);
    Assert(pPlayer != NULL);
    GamePlayer* pGamePlayer = (GamePlayer*)pPlayer;
    Assert(pGamePlayer!= NULL);
    if ( pGamePlayer->GetPlayerStatus() != PS_SERVER_NORMAL )
    {
        CacheGuid64Log(LOG_FILE_1, "ERROR CGTestHandler::pGamePlayer->GetPlayerStatus():GUID=%s.", Guid2StringObj(pGamePlayer->m_HumanGUID).GetString());
        return PACKET_EXE_CONTINUE;
    }
    Obj_Human* pHuman = pGamePlayer->GetHuman();
    Assert(pHuman);
    if ( !pGamePlayer->IsCanLogic() )
    {
        CacheGuid64Log(LOG_FILE_1, "ERROR CGTestHandler::pGamePlayer->IsCanLogic():GUID=%s Name:%s.", Guid2StringObj(pGamePlayer->m_HumanGUID).GetString(), pHuman->GetName());
        return PACKET_EXE_CONTINUE;
    }
    Scene* pScene = pHuman->getScene();
    if ( pScene == NULL )
    {
        Assert(FALSE);
        return PACKET_EXE_ERROR;
    }
    Assert(pScene->VerifyExecuteThread());
    //写下你的代码
    return PACKET_EXE_CONTINUE;
__LEAVE_FUNCTION
    return PACKET_EXE_ERROR;
}