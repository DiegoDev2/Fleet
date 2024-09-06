package main

// Hive - Hadoop-based data summarization, query, and analysis
// Homepage: https://hive.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installHive() {
	// Método 1: Descargar y extraer .tar.gz
	hive_tar_url := "https://www.apache.org/dyn/closer.lua?path=hive/hive-3.1.3/apache-hive-3.1.3-bin.tar.gz"
	hive_cmd_tar := exec.Command("curl", "-L", hive_tar_url, "-o", "package.tar.gz")
	err := hive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hive_zip_url := "https://www.apache.org/dyn/closer.lua?path=hive/hive-3.1.3/apache-hive-3.1.3-bin.zip"
	hive_cmd_zip := exec.Command("curl", "-L", hive_zip_url, "-o", "package.zip")
	err = hive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hive_bin_url := "https://www.apache.org/dyn/closer.lua?path=hive/hive-3.1.3/apache-hive-3.1.3-bin.bin"
	hive_cmd_bin := exec.Command("curl", "-L", hive_bin_url, "-o", "binary.bin")
	err = hive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hive_src_url := "https://www.apache.org/dyn/closer.lua?path=hive/hive-3.1.3/apache-hive-3.1.3-bin.src.tar.gz"
	hive_cmd_src := exec.Command("curl", "-L", hive_src_url, "-o", "source.tar.gz")
	err = hive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hive_cmd_direct := exec.Command("./binary")
	err = hive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hadoop")
exec.Command("latte", "install", "hadoop")
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
