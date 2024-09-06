package main

// Jasmin - Assembler for the Java Virtual Machine
// Homepage: https://jasmin.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installJasmin() {
	// Método 1: Descargar y extraer .tar.gz
	jasmin_tar_url := "https://downloads.sourceforge.net/project/jasmin/jasmin/jasmin-2.4/jasmin-2.4.zip"
	jasmin_cmd_tar := exec.Command("curl", "-L", jasmin_tar_url, "-o", "package.tar.gz")
	err := jasmin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jasmin_zip_url := "https://downloads.sourceforge.net/project/jasmin/jasmin/jasmin-2.4/jasmin-2.4.zip"
	jasmin_cmd_zip := exec.Command("curl", "-L", jasmin_zip_url, "-o", "package.zip")
	err = jasmin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jasmin_bin_url := "https://downloads.sourceforge.net/project/jasmin/jasmin/jasmin-2.4/jasmin-2.4.zip"
	jasmin_cmd_bin := exec.Command("curl", "-L", jasmin_bin_url, "-o", "binary.bin")
	err = jasmin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jasmin_src_url := "https://downloads.sourceforge.net/project/jasmin/jasmin/jasmin-2.4/jasmin-2.4.zip"
	jasmin_cmd_src := exec.Command("curl", "-L", jasmin_src_url, "-o", "source.tar.gz")
	err = jasmin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jasmin_cmd_direct := exec.Command("./binary")
	err = jasmin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
