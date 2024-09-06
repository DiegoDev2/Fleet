package main

// Owamp - Implementation of the One-Way Active Measurement Protocol
// Homepage: https://www.internet2.edu/products-services/performance-analytics/performance-tools/

import (
	"fmt"
	
	"os/exec"
)

func installOwamp() {
	// Método 1: Descargar y extraer .tar.gz
	owamp_tar_url := "https://software.internet2.edu/sources/owamp/owamp-3.4-10.tar.gz"
	owamp_cmd_tar := exec.Command("curl", "-L", owamp_tar_url, "-o", "package.tar.gz")
	err := owamp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	owamp_zip_url := "https://software.internet2.edu/sources/owamp/owamp-3.4-10.zip"
	owamp_cmd_zip := exec.Command("curl", "-L", owamp_zip_url, "-o", "package.zip")
	err = owamp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	owamp_bin_url := "https://software.internet2.edu/sources/owamp/owamp-3.4-10.bin"
	owamp_cmd_bin := exec.Command("curl", "-L", owamp_bin_url, "-o", "binary.bin")
	err = owamp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	owamp_src_url := "https://software.internet2.edu/sources/owamp/owamp-3.4-10.src.tar.gz"
	owamp_cmd_src := exec.Command("curl", "-L", owamp_src_url, "-o", "source.tar.gz")
	err = owamp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	owamp_cmd_direct := exec.Command("./binary")
	err = owamp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: i2util")
	exec.Command("latte", "install", "i2util").Run()
}
