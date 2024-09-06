package main

// Derby - Apache Derby is an embedded relational database running on JVM
// Homepage: https://db.apache.org/derby/

import (
	"fmt"
	
	"os/exec"
)

func installDerby() {
	// Método 1: Descargar y extraer .tar.gz
	derby_tar_url := "https://www.apache.org/dyn/closer.lua?path=db/derby/db-derby-10.17.1.0/db-derby-10.17.1.0-bin.tar.gz"
	derby_cmd_tar := exec.Command("curl", "-L", derby_tar_url, "-o", "package.tar.gz")
	err := derby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	derby_zip_url := "https://www.apache.org/dyn/closer.lua?path=db/derby/db-derby-10.17.1.0/db-derby-10.17.1.0-bin.zip"
	derby_cmd_zip := exec.Command("curl", "-L", derby_zip_url, "-o", "package.zip")
	err = derby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	derby_bin_url := "https://www.apache.org/dyn/closer.lua?path=db/derby/db-derby-10.17.1.0/db-derby-10.17.1.0-bin.bin"
	derby_cmd_bin := exec.Command("curl", "-L", derby_bin_url, "-o", "binary.bin")
	err = derby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	derby_src_url := "https://www.apache.org/dyn/closer.lua?path=db/derby/db-derby-10.17.1.0/db-derby-10.17.1.0-bin.src.tar.gz"
	derby_cmd_src := exec.Command("curl", "-L", derby_src_url, "-o", "source.tar.gz")
	err = derby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	derby_cmd_direct := exec.Command("./binary")
	err = derby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
