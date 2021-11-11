CGTest = class("CGTest", LuaRequestPacket)

function CGTest:Init()
    self:SetupGameServerPacket(2882)
end

function CGTest:WriteStream()
        index = WriteInt32(self.m_Score);
        index = WriteInt16(self.m_FlagCount);
        for i=1, m_FlagCount do
            index = WriteUInt64(self.m_Flag);
        end
        index = WriteInt16(self.m_Count);
        index = WriteInt16(self.m_NameCount);
        local Namelen = PS.PASocketStreamAssist.GetStringBytesLen(self.m_Name, false);
        index = WriteByte(Namelen);
        index = WriteCharArray(self.m_Name, Namelen, true, false);
        index = WriteInt16(self.m_Value4Count);
        index = WriteBytes(self.m_Value4, m_Value4Count);
        index = WriteSingle(self.m_wpPos.x);
        index = WriteSingle(self.m_wpPos.z);
        index = WriteSingle(self.m_Dir.x);
        index = WriteSingle(self.m_Dir.y);
        index = WriteSingle(self.m_Dir.z);
        index = WriteSingle(self.m_Pos.x);
        index = WriteSingle(self.m_Pos.y);
end