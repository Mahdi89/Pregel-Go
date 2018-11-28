module PingPong #(
parameter VERTEX_WIDTH = 64,
parameter EDGE_WIDTH = 32,
parameter EDGE_ADDRESS = 32,
parameter OUTPUT_WIDTH = 32
)(
input [VERTEX_WIDTH-1:0] 	Vertex_in,
input [EDGE_WIDTH-1:0] 		Edge_in,
output [EDGE_ADDRESS-1:0]	Edge_addr,
output [EDGE_WIDTH-1:0] 	Edge_out,
output [OUTPUT_WIDTH-1:0]	Update_val
);

reg [EDGE_ADDRESS-1:0] edge_addr;
reg [EDGE_WIDTH-1:0] edge_out;
reg [OUTPUT_WIDTH-1:0] update_val;

assign Edge_addr = edge_addr;
assign Edge_out = edge_out;
assign Update_val = update_val;

always @(Vertex_in)begin
	// Check the edges associated with a particular input vertex
	edge_addr <= Vertex_in[(VERTEX_WIDTH/2)-1:0];
	// TODO Sum on edge values associated with this vertex
	update_val <= Edge_in;
end


endmodule
