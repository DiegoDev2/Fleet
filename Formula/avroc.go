package main

// AvroC - Data serialization system
// Homepage: https://avro.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installAvroC() {
	// Método 1: Descargar y extraer .tar.gz
	avroc_tar_url := "https://github.com/apache/avro.git"
	avroc_cmd_tar := exec.Command("curl", "-L", avroc_tar_url, "-o", "package.tar.gz")
	err := avroc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avroc_zip_url := "https://github.com/apache/avro.git"
	avroc_cmd_zip := exec.Command("curl", "-L", avroc_zip_url, "-o", "package.zip")
	err = avroc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avroc_bin_url := "https://github.com/apache/avro.git"
	avroc_cmd_bin := exec.Command("curl", "-L", avroc_bin_url, "-o", "binary.bin")
	err = avroc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avroc_src_url := "https://github.com/apache/avro.git"
	avroc_cmd_src := exec.Command("curl", "-L", avroc_src_url, "-o", "source.tar.gz")
	err = avroc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avroc_cmd_direct := exec.Command("./binary")
	err = avroc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
