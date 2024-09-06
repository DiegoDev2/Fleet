package main

// ApacheFlink - Scalable batch and stream data processing
// Homepage: https://flink.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheFlink() {
	// Método 1: Descargar y extraer .tar.gz
	apacheflink_tar_url := "https://www.apache.org/dyn/closer.lua?path=flink/flink-1.20.0/flink-1.20.0-bin-scala_2.12.tgz"
	apacheflink_cmd_tar := exec.Command("curl", "-L", apacheflink_tar_url, "-o", "package.tar.gz")
	err := apacheflink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apacheflink_zip_url := "https://www.apache.org/dyn/closer.lua?path=flink/flink-1.20.0/flink-1.20.0-bin-scala_2.12.tgz"
	apacheflink_cmd_zip := exec.Command("curl", "-L", apacheflink_zip_url, "-o", "package.zip")
	err = apacheflink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apacheflink_bin_url := "https://www.apache.org/dyn/closer.lua?path=flink/flink-1.20.0/flink-1.20.0-bin-scala_2.12.tgz"
	apacheflink_cmd_bin := exec.Command("curl", "-L", apacheflink_bin_url, "-o", "binary.bin")
	err = apacheflink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apacheflink_src_url := "https://www.apache.org/dyn/closer.lua?path=flink/flink-1.20.0/flink-1.20.0-bin-scala_2.12.tgz"
	apacheflink_cmd_src := exec.Command("curl", "-L", apacheflink_src_url, "-o", "source.tar.gz")
	err = apacheflink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apacheflink_cmd_direct := exec.Command("./binary")
	err = apacheflink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
