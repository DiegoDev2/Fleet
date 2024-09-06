package main

// Kettle - Pentaho Data Integration software
// Homepage: https://www.hitachivantara.com/en-us/products/pentaho-plus-platform/data-integration-analytics.html

import (
	"fmt"
	
	"os/exec"
)

func installKettle() {
	// Método 1: Descargar y extraer .tar.gz
	kettle_tar_url := "https://privatefilesbucket-community-edition.s3.us-west-2.amazonaws.com/9.4.0.0-343/ce/client-tools/pdi-ce-9.4.0.0-343.zip"
	kettle_cmd_tar := exec.Command("curl", "-L", kettle_tar_url, "-o", "package.tar.gz")
	err := kettle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kettle_zip_url := "https://privatefilesbucket-community-edition.s3.us-west-2.amazonaws.com/9.4.0.0-343/ce/client-tools/pdi-ce-9.4.0.0-343.zip"
	kettle_cmd_zip := exec.Command("curl", "-L", kettle_zip_url, "-o", "package.zip")
	err = kettle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kettle_bin_url := "https://privatefilesbucket-community-edition.s3.us-west-2.amazonaws.com/9.4.0.0-343/ce/client-tools/pdi-ce-9.4.0.0-343.zip"
	kettle_cmd_bin := exec.Command("curl", "-L", kettle_bin_url, "-o", "binary.bin")
	err = kettle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kettle_src_url := "https://privatefilesbucket-community-edition.s3.us-west-2.amazonaws.com/9.4.0.0-343/ce/client-tools/pdi-ce-9.4.0.0-343.zip"
	kettle_cmd_src := exec.Command("curl", "-L", kettle_src_url, "-o", "source.tar.gz")
	err = kettle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kettle_cmd_direct := exec.Command("./binary")
	err = kettle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
