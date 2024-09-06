package main

// Hadoop - Framework for distributed processing of large data sets
// Homepage: https://hadoop.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installHadoop() {
	// Método 1: Descargar y extraer .tar.gz
	hadoop_tar_url := "https://www.apache.org/dyn/closer.lua?path=hadoop/common/hadoop-3.4.0/hadoop-3.4.0.tar.gz"
	hadoop_cmd_tar := exec.Command("curl", "-L", hadoop_tar_url, "-o", "package.tar.gz")
	err := hadoop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hadoop_zip_url := "https://www.apache.org/dyn/closer.lua?path=hadoop/common/hadoop-3.4.0/hadoop-3.4.0.zip"
	hadoop_cmd_zip := exec.Command("curl", "-L", hadoop_zip_url, "-o", "package.zip")
	err = hadoop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hadoop_bin_url := "https://www.apache.org/dyn/closer.lua?path=hadoop/common/hadoop-3.4.0/hadoop-3.4.0.bin"
	hadoop_cmd_bin := exec.Command("curl", "-L", hadoop_bin_url, "-o", "binary.bin")
	err = hadoop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hadoop_src_url := "https://www.apache.org/dyn/closer.lua?path=hadoop/common/hadoop-3.4.0/hadoop-3.4.0.src.tar.gz"
	hadoop_cmd_src := exec.Command("curl", "-L", hadoop_src_url, "-o", "source.tar.gz")
	err = hadoop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hadoop_cmd_direct := exec.Command("./binary")
	err = hadoop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
