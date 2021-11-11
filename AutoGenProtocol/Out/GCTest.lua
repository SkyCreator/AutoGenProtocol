local GCTest = class("GCTest", LuaResponsePacket)
function GCTest:ReadStream()
        self.m_Score = self.ReadInt32();
        self.m_FlagCount = self.ReadInt16();
        self.m_Flag = {}
        for i=1, m_FlagCount do
            self.m_Flag[i] = self.ReadUInt64();
        end
        self.m_Count = self.ReadInt16();
        self.m_Value4Count = self.ReadInt16();
        self.m_Value4 = self.ReadBytes(self.m_Value4, m_Value4Count);
        local wpPosx = self.ReadSingle();
        local wpPosz = self.ReadSingle();
        self.wpPos = {x=wpPosx, z=wpPosz}
        local Dirx = self.ReadSingle();
        local Diry = self.ReadSingle();
        local Dirz = self.ReadSingle();
        self.Dir = Vector3(Dirx, Diry, Dirz)
        self.m_PosCount = self.ReadInt16();
        self.m_Pos = {}
        for i=1, m_PosCount do
            local Posx = self.ReadSingle();
            local Posy = self.ReadSingle();
            self.Pos[i] = Vector2(Posx, Posy)
        end
        self.m_TitleCount = self.ReadInt16();
        self.m_Title = self.ReadString(self.m_TitleCount);
end

function GCTest:Handler()
end
return GCTest
