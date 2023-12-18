%%change this variable, to change range of x-axis for plots 2 and 3. 
%(depending on efficiency of used VM, 0-500 does not make any sense)
xaxis = 250; %paper uses 500


%%change wich figure should be displayed
%(since online editors can only visualize one figure, we did choose this option)
% 0 = all figures (e.g. if you run it via full octave)
% 1 = Figure 8: encrypting one record
% 2 = Figure 9: generating decryption key
% 3 = Figure 10: evaluating function
fig = 1; 

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%Please cahnge to results produced by experiments!
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%


DiffPIPE_attr5 = [
5, 11,1,34.333333333333336,111; 
10, 25.5,1,63,211; 
20, 59.666666666666664,1.3333333333333333,135.66666666666666,398.3333333333333; 
30, 81.66666666666667,2.3333333333333335,179.33333333333334,539; 
];

DiffPIPE_attr10 = [
5, 29.333333333333332,3,58.666666666666664,164.66666666666666; 
10, 89,3,114.66666666666667,333; 
20, 190,7,223.5,719.5; 
30, 233.33333333333334,2.3333333333333335,301,880.3333333333334; 
];

DiffPIPE_attr20 = [
5, 131,4.666666666666667,85.33333333333333,242.33333333333334; 
10, 232.33333333333334,6.666666666666667,156.66666666666666,458; 
20, 542.6666666666666,4,325.3333333333333,1023.6666666666666; 
30, 652,4,416,1293.6666666666667; 
];

DiffPIPE_attr30 = [
5, 277,7,109.66666666666667,308.6666666666667; 
10, 517.3333333333334,6.333333333333333,194.33333333333334,568.3333333333334; 
20, 920.3333333333334,6.333333333333333,402,1114.3333333333333; 
30, 1235.6666666666667,5,520.6666666666666,1685.6666666666667; 
];
noisyDOT_attr5 = [
5, 31.666666666666668,6,42.333333333333336,128.66666666666666; 
10, 64,2.3333333333333335,96.33333333333333,266.3333333333333; 
20, 133,2,159.5,499.5; 
30, 168,4.333333333333333,228.66666666666666,729.3333333333334; 
];

noisyDOT_attr10 = [
5, 98.5,4,61,211.5; 
10, 252.33333333333334,3.6666666666666665,150.33333333333334,516.3333333333334; 
20, 539.3333333333334,4.333333333333333,334.6666666666667,1089.6666666666667; 
30, 650.5,4,442,1358.5; 
];

noisyDOT_attr20 = [
5, 511.3333333333333,8,127,406; 
10, 972.5,13,249.5,824.5; 
20, 1775.6666666666667,8,480,1725; 
30, 2517.3333333333335,7.666666666666667,734.6666666666666,3844; 
];

noisyDOT_attr30 = [
5, 1202,11,168,520; 
10, 2088.3333333333335,11.666666666666666,313.6666666666667,1077; 
20, 3839.6666666666665,12.666666666666666,636.3333333333334,2100.3333333333335; 
30, 5545.333333333333,10.333333333333334,933.3333333333334,3114; 
];





%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%


%eval:
x = DiffPIPE_attr10(:,1);
ot_eval_10 = DiffPIPE_attr10(:, 5)/1000;
ot_eval_20 = DiffPIPE_attr20(:, 5)/1000;
ot_eval_30 = DiffPIPE_attr30(:, 5)/1000;

nh_eval_10 = noisyDOT_attr10(:, 5)/1000;
nh_eval_20 = noisyDOT_attr20(:, 5)/1000;
nh_eval_30 = noisyDOT_attr30(:, 5)/1000;



%keygen
ot_kg_10 = DiffPIPE_attr10(:, 4)/1000;
ot_kg_20 = DiffPIPE_attr20(:, 4)/1000;
ot_kg_30 = DiffPIPE_attr30(:, 4)/1000;

nh_kg_10 = noisyDOT_attr10(:, 4)/1000;
nh_kg_20 = noisyDOT_attr20(:, 4)/1000;
nh_kg_30 = noisyDOT_attr30(:, 4)/1000;



%enc
data_nh_enc = [5, noisyDOT_attr5(1,3);
10, noisyDOT_attr10(1,3);
20, noisyDOT_attr20(1,3);
30, noisyDOT_attr30(1,3);
];
data_ot_enc = [5, DiffPIPE_attr5(1,3);
10, DiffPIPE_attr10(1,3);
20, DiffPIPE_attr20(1,3);
30, DiffPIPE_attr30(1, 3);
];

xE = data_nh_enc(:,1);
nh_enc = data_nh_enc(:,2);
ot_enc = data_ot_enc(:,2);


%lin reg enc
lr_nh_enc = polyfit(xE, nh_enc,1);
lr_ot_enc = polyfit(xE, ot_enc,1);

%lin reg eval
lr_ot_eval_10 = polyfit(x,ot_eval_10,1);
lr_ot_eval_20 = polyfit(x,ot_eval_20,1);
lr_ot_eval_30 = polyfit(x,ot_eval_30,1);

lr_nh_eval_10 = polyfit(x,nh_eval_10,1);
lr_nh_eval_20 = polyfit(x,nh_eval_20,1);
lr_nh_eval_30 = polyfit(x,nh_eval_30,1);

%lin reg kg
lr_ot_kg_10 = polyfit(x,ot_kg_10,1);
lr_ot_kg_20 = polyfit(x,ot_kg_20,1);
lr_ot_kg_30 = polyfit(x,ot_kg_30,1);

lr_nh_kg_10 = polyfit(x,nh_kg_10,1);
lr_nh_kg_20 = polyfit(x,nh_kg_20,1);
lr_nh_kg_30 = polyfit(x,nh_kg_30,1);




if (fig == 0)
  figure(1)
  

  y = 1:50;
  plot(xE, ot_enc, "xk");
  hold;
  plot( lr_ot_enc(1)*y + lr_ot_enc(2), "--k");
  plot(xE, nh_enc, "xk");
  plot( lr_nh_enc(1)*y + lr_nh_enc(2), "k");

  legend("DiffPIPE", "", "nDOT", "");
  xlabel("#attritbutes");
    title("Time needed to encrypt one record");
  ylabel("time in millisec");

  set(gca,  "fontsize", 15);
  figure(2) 
  y = 1:xaxis;
  plot(x, ot_eval_10, "xk");
  hold;
  plot( lr_ot_eval_10(1)*y + lr_ot_eval_10(2), "--k");
  plot(x, ot_eval_30, "xm");
  plot( lr_ot_eval_30(1)*y + lr_ot_eval_30(2), "--m");
  plot(x, nh_eval_10, "xk");
  plot( lr_nh_eval_10(1)*y + lr_nh_eval_10(2), "k");
  plot(x, nh_eval_30, "xm");  
  plot( lr_nh_eval_30(1)*y + lr_nh_eval_30(2), "m");
  legend("10 attr DiffPIPE", "", "30 attr DiffPIPE", "", "10 attr nDOT", "", "30 attr nDOT", "" )
  title("Time for evaluating function");
  xlabel("#records");
  ylabel("time in sec");
  set(gca,  "fontsize", 15);

  figure(3) 
  y = 1:xaxis;
  plot(x, ot_kg_10, "xk");
  hold;
  plot( lr_ot_kg_10(1)*y + lr_ot_kg_10(2), "--k");
  plot(x, ot_kg_30, "xm");
  plot( lr_ot_kg_30(1)*y + lr_ot_kg_30(2), "--m");
  plot(x, nh_kg_10, "xk");
  plot( lr_nh_kg_10(1)*y + lr_nh_kg_10(2), "k");
  plot(x, nh_kg_30, "xm");
  plot( lr_nh_kg_30(1)*y + lr_nh_kg_30(2), "m");
  legend("10 attr DiffPIPE", "", "30 attr DiffPIPE", "", "10 attr nDOT", "", "30 attr nDOT", "" )
  title("Time for generating decryption key");
  xlabel("#records");
  ylabel("time in sec");
  
set(gca,  "fontsize", 15);

endif

if (fig == 1)

  y = 1:50;
  plot(xE, ot_enc, "xk");
  hold;
  plot( lr_ot_enc(1)*y + lr_ot_enc(2), "--k");
  plot(xE, nh_enc, "xk");
  plot( lr_nh_enc(1)*y + lr_nh_enc(2), "k");

  legend("DiffPIPE", "", "nDOT", "");
  xlabel("#attritbutes");
    title("Time needed to encrypt one record");
  ylabel("time in millisec");
  
set(gca,  "fontsize", 15);
endif

if (fig == 2)
  y = 1:xaxis;
  plot(x, ot_eval_10, "xk");
  hold;
  plot( lr_ot_eval_10(1)*y + lr_ot_eval_10(2), "--k");
  plot(x, ot_eval_30, "xm");
  plot( lr_ot_eval_30(1)*y + lr_ot_eval_30(2), "--m");
  plot(x, nh_eval_10, "xk");
  plot( lr_nh_eval_10(1)*y + lr_nh_eval_10(2), "k");
  plot(x, nh_eval_30, "xm");  
  plot( lr_nh_eval_30(1)*y + lr_nh_eval_30(2), "m");
  legend("10 attr DiffPIPE", "", "30 attr DiffPIPE", "", "10 attr nDOT", "", "30 attr nDOT", "" )
  title("Time for evaluating function");
  xlabel("#records");
  ylabel("time in sec");
  
set(gca,  "fontsize", 15);
endif

if (fig == 3)

  y = 1:xaxis;
  plot(x, ot_kg_10, "xk");
  hold;
  plot( lr_ot_kg_10(1)*y + lr_ot_kg_10(2), "--k");
  plot(x, ot_kg_30, "xm");
  plot( lr_ot_kg_30(1)*y + lr_ot_kg_30(2), "--m");
  plot(x, nh_kg_10, "xk");
  plot( lr_nh_kg_10(1)*y + lr_nh_kg_10(2), "k");
  plot(x, nh_kg_30, "xm");
  plot( lr_nh_kg_30(1)*y + lr_nh_kg_30(2), "m");
  legend("10 attr DiffPIPE", "", "30 attr DiffPIPE", "", "10 attr nDOT", "", "30 attr nDOT", "" )
  title("Time for generating decryption key");
  xlabel("#records");
  ylabel("time in sec");

set(gca,  "fontsize", 15);
endif



