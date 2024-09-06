package main

// Librdkafka - Apache Kafka C/C++ library
// Homepage: https://github.com/confluentinc/librdkafka

import (
	"fmt"
	
	"os/exec"
)

func installLibrdkafka() {
	// Método 1: Descargar y extraer .tar.gz
	librdkafka_tar_url := "https://github.com/confluentinc/librdkafka/archive/refs/tags/v2.5.3.tar.gz"
	librdkafka_cmd_tar := exec.Command("curl", "-L", librdkafka_tar_url, "-o", "package.tar.gz")
	err := librdkafka_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librdkafka_zip_url := "https://github.com/confluentinc/librdkafka/archive/refs/tags/v2.5.3.zip"
	librdkafka_cmd_zip := exec.Command("curl", "-L", librdkafka_zip_url, "-o", "package.zip")
	err = librdkafka_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librdkafka_bin_url := "https://github.com/confluentinc/librdkafka/archive/refs/tags/v2.5.3.bin"
	librdkafka_cmd_bin := exec.Command("curl", "-L", librdkafka_bin_url, "-o", "binary.bin")
	err = librdkafka_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librdkafka_src_url := "https://github.com/confluentinc/librdkafka/archive/refs/tags/v2.5.3.src.tar.gz"
	librdkafka_cmd_src := exec.Command("curl", "-L", librdkafka_src_url, "-o", "source.tar.gz")
	err = librdkafka_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librdkafka_cmd_direct := exec.Command("./binary")
	err = librdkafka_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: lzlib")
exec.Command("latte", "install", "lzlib")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
