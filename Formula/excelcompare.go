package main

// ExcelCompare - Command-line tool (and API) for diffing Excel Workbooks
// Homepage: https://github.com/na-ka-na/ExcelCompare

import (
	"fmt"
	
	"os/exec"
)

func installExcelCompare() {
	// Método 1: Descargar y extraer .tar.gz
	excelcompare_tar_url := "https://github.com/na-ka-na/ExcelCompare/releases/download/0.7.0/ExcelCompare-0.7.0.zip"
	excelcompare_cmd_tar := exec.Command("curl", "-L", excelcompare_tar_url, "-o", "package.tar.gz")
	err := excelcompare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	excelcompare_zip_url := "https://github.com/na-ka-na/ExcelCompare/releases/download/0.7.0/ExcelCompare-0.7.0.zip"
	excelcompare_cmd_zip := exec.Command("curl", "-L", excelcompare_zip_url, "-o", "package.zip")
	err = excelcompare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	excelcompare_bin_url := "https://github.com/na-ka-na/ExcelCompare/releases/download/0.7.0/ExcelCompare-0.7.0.zip"
	excelcompare_cmd_bin := exec.Command("curl", "-L", excelcompare_bin_url, "-o", "binary.bin")
	err = excelcompare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	excelcompare_src_url := "https://github.com/na-ka-na/ExcelCompare/releases/download/0.7.0/ExcelCompare-0.7.0.zip"
	excelcompare_cmd_src := exec.Command("curl", "-L", excelcompare_src_url, "-o", "source.tar.gz")
	err = excelcompare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	excelcompare_cmd_direct := exec.Command("./binary")
	err = excelcompare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
