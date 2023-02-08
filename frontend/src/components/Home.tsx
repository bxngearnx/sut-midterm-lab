// import { Link as RouterLink } from "react-router-dom";
// import Button from "@mui/material/Button";
// import Grid from "@mui/material/Grid";
// import Container from "@mui/material/Container";
// import Paper from "@mui/material/Paper";
// import Box from "@mui/material/Box";
// import {Homebar} from "./Homebar";
// import DataAdmin from "./DataAdmin";
// import React from "react";
// import { CssBaseline } from "@mui/material";

// function Home() {
//   return (
//     <div className="Home" id="outer-container">
//       <Homebar pageWrapId={"page-Home"} outerContainerId={"outer-container"} />
//       <div id="page-Home">
//         <React.Fragment>
//           <CssBaseline />
//           <Container maxWidth="lg">
//             <Paper sx={{ padding: 2 }}>
//               <Box display={"flex"}>
//                 <Box sx={{ flexGrow: 1 }}>
//                   <Grid container spacing={2}>
//                     <Grid item xs={2}>
//                       <Button
//                         variant="contained"
//                         size="large"
//                         fullWidth
//                         component={RouterLink}
//                         to="/DataAdmin"
//                       >
//                         Dataadmin
//                       </Button>
//                     </Grid>
//                     <Grid item xs={2}>
//                       <Button
//                         variant="contained"
//                         size="large"
//                         fullWidth
//                         component={RouterLink}
//                         to="/DataPostponement"
//                       >
//                         Datapostponement
//                       </Button>
//                     </Grid>
//                   </Grid>
//                 </Box>
//               </Box>
//             </Paper>
//           </Container>
//         </React.Fragment>
//       </div>
//     </div>
//   );
// }

// export default Home;
import Grid from "@mui/material/Grid";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { CssBaseline } from "@mui/material";
import React from "react";

import { Homebar } from "./Bar-Home";

function Home() {
  return (
    <div className="Home" id="outer-container">
      <Homebar pageWrapId={"page-Home"} outerContainerId={"outer-container"} />
      <div id="page-Home">
        <React.Fragment>
          <CssBaseline />
          <Grid container component="main" sx={{ height: "100vh" }}>
            <Grid
              item
              xs={false}
              sm={4}
              md={12}
              sx={{
                backgroundImage:
                  "url(https://images.unsplash.com/photo-1571235854001-2c64e3fdd06a?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80)",
                backgroundRepeat: "secondary",
                backgroundColor: (t) =>
                  t.palette.mode === "dark"
                    ? t.palette.grey[50]
                    : t.palette.grey[900],
                backgroundSize: "cover",
                backgroundPosition: "center",
              }}
            >
              <div>
                <Container>
                  <Box>
                    <h1 style={{ textAlign: "center", color: "#ffffff" }}>
                      Test
                    </h1>
                  </Box>
                </Container>
              </div>
            </Grid>
          </Grid>
        </React.Fragment>
      </div>
    </div>
  );
}

export default Home;