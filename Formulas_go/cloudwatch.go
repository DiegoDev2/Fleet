package main

// CloudWatch - Amazon CloudWatch command-line Tool
// Homepage: https://aws.amazon.com/developertools/2534

import (
	"fmt"
	
	"os/exec"
)

func installCloudWatch() {
	// Método 1: Descargar y extraer .tar.gz
	cloudwatch_tar_url := "https://ec2-downloads.s3.amazonaws.com/CloudWatch-2010-08-01.zip"
	cloudwatch_cmd_tar := exec.Command("curl", "-L", cloudwatch_tar_url, "-o", "package.tar.gz")
	err := cloudwatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudwatch_zip_url := "https://ec2-downloads.s3.amazonaws.com/CloudWatch-2010-08-01.zip"
	cloudwatch_cmd_zip := exec.Command("curl", "-L", cloudwatch_zip_url, "-o", "package.zip")
	err = cloudwatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudwatch_bin_url := "https://ec2-downloads.s3.amazonaws.com/CloudWatch-2010-08-01.zip"
	cloudwatch_cmd_bin := exec.Command("curl", "-L", cloudwatch_bin_url, "-o", "binary.bin")
	err = cloudwatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudwatch_src_url := "https://ec2-downloads.s3.amazonaws.com/CloudWatch-2010-08-01.zip"
	cloudwatch_cmd_src := exec.Command("curl", "-L", cloudwatch_src_url, "-o", "source.tar.gz")
	err = cloudwatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudwatch_cmd_direct := exec.Command("./binary")
	err = cloudwatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
