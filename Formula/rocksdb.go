package main

// Rocksdb - Embeddable, persistent key-value store for fast storage
// Homepage: https://rocksdb.org/

import (
	"fmt"
	
	"os/exec"
)

func installRocksdb() {
	// Método 1: Descargar y extraer .tar.gz
	rocksdb_tar_url := "https://github.com/facebook/rocksdb/archive/refs/tags/v9.5.2.tar.gz"
	rocksdb_cmd_tar := exec.Command("curl", "-L", rocksdb_tar_url, "-o", "package.tar.gz")
	err := rocksdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rocksdb_zip_url := "https://github.com/facebook/rocksdb/archive/refs/tags/v9.5.2.zip"
	rocksdb_cmd_zip := exec.Command("curl", "-L", rocksdb_zip_url, "-o", "package.zip")
	err = rocksdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rocksdb_bin_url := "https://github.com/facebook/rocksdb/archive/refs/tags/v9.5.2.bin"
	rocksdb_cmd_bin := exec.Command("curl", "-L", rocksdb_bin_url, "-o", "binary.bin")
	err = rocksdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rocksdb_src_url := "https://github.com/facebook/rocksdb/archive/refs/tags/v9.5.2.src.tar.gz"
	rocksdb_cmd_src := exec.Command("curl", "-L", rocksdb_src_url, "-o", "source.tar.gz")
	err = rocksdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rocksdb_cmd_direct := exec.Command("./binary")
	err = rocksdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
