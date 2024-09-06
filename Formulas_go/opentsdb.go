package main

// Opentsdb - Scalable, distributed Time Series Database
// Homepage: http://opentsdb.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpentsdb() {
	// Método 1: Descargar y extraer .tar.gz
	opentsdb_tar_url := "https://github.com/OpenTSDB/opentsdb/archive/refs/tags/v2.4.1.tar.gz"
	opentsdb_cmd_tar := exec.Command("curl", "-L", opentsdb_tar_url, "-o", "package.tar.gz")
	err := opentsdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opentsdb_zip_url := "https://github.com/OpenTSDB/opentsdb/archive/refs/tags/v2.4.1.zip"
	opentsdb_cmd_zip := exec.Command("curl", "-L", opentsdb_zip_url, "-o", "package.zip")
	err = opentsdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opentsdb_bin_url := "https://github.com/OpenTSDB/opentsdb/archive/refs/tags/v2.4.1.bin"
	opentsdb_cmd_bin := exec.Command("curl", "-L", opentsdb_bin_url, "-o", "binary.bin")
	err = opentsdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opentsdb_src_url := "https://github.com/OpenTSDB/opentsdb/archive/refs/tags/v2.4.1.src.tar.gz"
	opentsdb_cmd_src := exec.Command("curl", "-L", opentsdb_src_url, "-o", "source.tar.gz")
	err = opentsdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opentsdb_cmd_direct := exec.Command("./binary")
	err = opentsdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gnuplot")
exec.Command("latte", "install", "gnuplot")
	fmt.Println("Instalando dependencia: hbase")
exec.Command("latte", "install", "hbase")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
