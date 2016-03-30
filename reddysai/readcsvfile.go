package main 
import ("fmt"; "encoding/csv"; "os")

type TestRecord struct {
	policyID string
	statecode string
	country string
	eqsiteLimit string
	husiteLimit string
	flsiteLimit string
	frsiteLimit string
	tiv_2011 string
	tiv_2012 string
	eqsitedeductible string
	husitedeductible string
	flsitedeductible string
	frsitedeductible string
	point_latitude string
	point_langitude string
	line string
	construction string
	point_granularity string
}

func main() {
	csvfile, err := os.Open("/home/agiratech/Downloads/myfile_sample.csv")
	if err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	//display the standard output
	for _, each := range rawCSVdata {
			fmt.Printf("%s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\t %s\n",each[0],each[1],each[2],each[3],each[4],each[5],each[6],each[7],each[8],each[9],each[10],each[11],each[12],each[13],each[14],each[15],each[16],each[17])
	}
	//now store the date into the struct
	var onerecord TestRecord
	var allrecords []TestRecord

	for _,each := range rawCSVdata {
		onerecord.policyID=each[0]
		onerecord.statecode=each[1]
		onerecord.country=each[2]
		onerecord.eqsiteLimit=each[3]
		onerecord.husiteLimit=each[4]
		onerecord.flsiteLimit=each[5]
		onerecord.frsiteLimit=each[6]
		onerecord.tiv_2011=each[7]
		onerecord.tiv_2012=each[8]
		onerecord.eqsitedeductible=each[9]
		onerecord.husitedeductible=each[10]
		onerecord.flsitedeductible=each[11]
		onerecord.frsitedeductible=each[12]
		onerecord.point_latitude=each[13]
		onerecord.point_langitude=each[14]
		onerecord.line=each[15]
		onerecord.construction=each[16]
		onerecord.point_granularity=each[17]
		
		allrecords =append(allrecords, onerecord)

	}
	fmt.Println(allrecords)
}