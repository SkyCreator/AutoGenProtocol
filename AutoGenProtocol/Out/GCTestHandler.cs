using GameFramework.Network;

namespace ProjectS
{
    public class GCTestHandler : PacketHandlerBase
    {
        public override int Id
        {
            get
            {
                return 2883;
            }
        }

        public override void Handle(object sender, Packet packet)
        {
            base.Handle(sender, packet);
        }
    }
}
