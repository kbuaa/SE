import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page,pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDoctor } from '../../api/models/EntDoctor'; // import interface User
import { EntPatient } from '../../api/models/EntPatient'; // import interface Car
import { EntDrug } from '../../api/models/EntDrug'; // import interface job
import { EntNurse } from '../../api/models/EntNurse'; // import interface job
// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));


const Prescription: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

 
  const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [drugs, setDrugs] = React.useState<EntDrug[]>([]);
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);
    
  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getDoctors = async () => {
    const res = await http.listDoctor({ limit: 2, offset: 0 });
    setDoctors(res);
  };

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 3, offset: 0 });
    setPatients(res);
  };

  const getDrug = async () => {
    const res = await http.listDrug({ limit: 7, offset: 0 });
    setDrugs(res);
  };

  const getNurse = async () => {
    const res = await http.listNurse({ limit: 2, offset: 0 });
    setNurses(res);
  };

  const PrescriptionNotehandleChange = (event: any) => {
    setPrescriptionNote(event.target.value as string);
  };
  const PrescriptiontDateTimehandleChange = (event: any) => {
    setPrescriptionDateTime(event.target.value as string);
  };
  const DoctorhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDoctorID(event.target.value as number);
  };
  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientID(event.target.value as number);
  };
  const DrughandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDrugID(event.target.value as number);
  };

  const NursehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNurseID(event.target.value as number);
  };

  const [prescriptionnote, setPrescriptionNote] = React.useState(String);
  const [added, setPrescriptionDateTime] = React.useState(String);
  const [doctorID, setDoctorID] = React.useState(Number);
  const [patientID, setPatientID] = React.useState(Number);
  const [drugID, setDrugID] = React.useState(Number);
  const [nurseID, setNurseID] = React.useState(Number);



  // Lifecycle Hooks
  useEffect(() => {
    getDoctors();
    getPatient();
    getDrug();
    getNurse();
  }, []);


  // function save data
  function save() {
    const prescriptions = {
      note: prescriptionnote,
      added: added + ":00+00:00",
      doctor: doctorID,
      patient: patientID,
      drug: drugID,
      nurse: nurseID,
    }
    const apiUrl = 'http://localhost:8080/api/v1/prescriptions';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(prescriptions),
    };
    console.log(prescriptions); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
        
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }

      });
}
  
  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Prescription System`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>ผู้ป่วย</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="patient"
                  value={patientID}
                  onChange={PatienthandleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>แพทย์ผู้สั่งยา</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="doctor"
                  value={doctorID}
                  onChange={DoctorhandleChange}
                >
                  {doctors.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.doctorName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ยาที่สั่ง</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="drug"
                  value={drugID}
                  onChange={DrughandleChange}
                >
                  {drugs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.drugName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>วันเวลาที่สั่งยา</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันเวลาที่สั่งยา"
                  name="added"
                  type="datetime-local"
                  value={added}
                  onChange={PrescriptiontDateTimehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>พยาบาลผู้บันทึก</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="nurse"
                  value={nurseID}
                  onChange={NursehandleChange}
                >
                  {nurses.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nurseName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>หมายเหตุ</div>
            </Grid>
            <Grid item xs={8}>
            <TextField
               name="note"
               label=""
               multiline
               defaultValue="Default Value"
               variant="outlined"
               value={prescriptionnote} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={PrescriptionNotehandleChange}
              />
            </Grid>
            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
            <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                SAVE
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Prescription;