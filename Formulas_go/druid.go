package main

// Druid - High-performance, column-oriented, distributed data store
// Homepage: https://druid.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installDruid() {
	// Método 1: Descargar y extraer .tar.gz
	druid_tar_url := "https://dlcdn.apache.org/druid/30.0.0/apache-druid-30.0.0-bin.tar.gz"
	druid_cmd_tar := exec.Command("curl", "-L", druid_tar_url, "-o", "package.tar.gz")
	err := druid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	druid_zip_url := "https://dlcdn.apache.org/druid/30.0.0/apache-druid-30.0.0-bin.zip"
	druid_cmd_zip := exec.Command("curl", "-L", druid_zip_url, "-o", "package.zip")
	err = druid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	druid_bin_url := "https://dlcdn.apache.org/druid/30.0.0/apache-druid-30.0.0-bin.bin"
	druid_cmd_bin := exec.Command("curl", "-L", druid_bin_url, "-o", "binary.bin")
	err = druid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	druid_src_url := "https://dlcdn.apache.org/druid/30.0.0/apache-druid-30.0.0-bin.src.tar.gz"
	druid_cmd_src := exec.Command("curl", "-L", druid_src_url, "-o", "source.tar.gz")
	err = druid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	druid_cmd_direct := exec.Command("./binary")
	err = druid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zookeeper")
exec.Command("latte", "install", "zookeeper")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
