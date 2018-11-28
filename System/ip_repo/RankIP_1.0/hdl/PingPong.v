module PingPong #(
parameter VERTEX_WIDTH = 64,
parameter EDGE_WIDTH = 32,
parameter EDGE_ADDRESS = 32,
parameter OUTPUT_WIDTH = 32
)(
input [VERTEX_WIDTH-1:0] 	Vertex_in,
input [EDGE_WIDTH-1:0] 		Edge_in,
output reg [EDGE_ADDRESS-1:0]	Edge_addr,
output reg [EDGE_WIDTH-1:0] 	Edge_out,
output reg [OUTPUT_WIDTH-1:0]	Update_val
);


always @(Vertex_in)begin
	// Check the edges associated with a particular input vertex
	Edge_addr <= Vertex_in[(VERTEX_WIDTH/2)-1:0];
	// TODO Sum on edge values associated with this vertex
	Update_val <= Edge_in;
end


endmodule
