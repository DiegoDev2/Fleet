package main

// Leveldb - Key-value storage library with ordered mapping
// Homepage: https://github.com/google/leveldb/

import (
	"fmt"
	
	"os/exec"
)

func installLeveldb() {
	// Método 1: Descargar y extraer .tar.gz
	leveldb_tar_url := "https://github.com/google/leveldb/archive/refs/tags/1.23.tar.gz"
	leveldb_cmd_tar := exec.Command("curl", "-L", leveldb_tar_url, "-o", "package.tar.gz")
	err := leveldb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leveldb_zip_url := "https://github.com/google/leveldb/archive/refs/tags/1.23.zip"
	leveldb_cmd_zip := exec.Command("curl", "-L", leveldb_zip_url, "-o", "package.zip")
	err = leveldb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leveldb_bin_url := "https://github.com/google/leveldb/archive/refs/tags/1.23.bin"
	leveldb_cmd_bin := exec.Command("curl", "-L", leveldb_bin_url, "-o", "binary.bin")
	err = leveldb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leveldb_src_url := "https://github.com/google/leveldb/archive/refs/tags/1.23.src.tar.gz"
	leveldb_cmd_src := exec.Command("curl", "-L", leveldb_src_url, "-o", "source.tar.gz")
	err = leveldb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leveldb_cmd_direct := exec.Command("./binary")
	err = leveldb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gperftools")
	exec.Command("latte", "install", "gperftools").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
}
