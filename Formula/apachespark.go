package main

// ApacheSpark - Engine for large-scale data processing
// Homepage: https://spark.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheSpark() {
	// Método 1: Descargar y extraer .tar.gz
	apachespark_tar_url := "https://dlcdn.apache.org/spark/spark-3.5.2/spark-3.5.2-bin-hadoop3.tgz"
	apachespark_cmd_tar := exec.Command("curl", "-L", apachespark_tar_url, "-o", "package.tar.gz")
	err := apachespark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachespark_zip_url := "https://dlcdn.apache.org/spark/spark-3.5.2/spark-3.5.2-bin-hadoop3.tgz"
	apachespark_cmd_zip := exec.Command("curl", "-L", apachespark_zip_url, "-o", "package.zip")
	err = apachespark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachespark_bin_url := "https://dlcdn.apache.org/spark/spark-3.5.2/spark-3.5.2-bin-hadoop3.tgz"
	apachespark_cmd_bin := exec.Command("curl", "-L", apachespark_bin_url, "-o", "binary.bin")
	err = apachespark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachespark_src_url := "https://dlcdn.apache.org/spark/spark-3.5.2/spark-3.5.2-bin-hadoop3.tgz"
	apachespark_cmd_src := exec.Command("curl", "-L", apachespark_src_url, "-o", "source.tar.gz")
	err = apachespark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachespark_cmd_direct := exec.Command("./binary")
	err = apachespark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}
