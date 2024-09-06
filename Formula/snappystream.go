package main

// Snappystream - C++ snappy stream realization (compatible with snappy)
// Homepage: https://github.com/hoxnox/snappystream

import (
	"fmt"
	
	"os/exec"
)

func installSnappystream() {
	// Método 1: Descargar y extraer .tar.gz
	snappystream_tar_url := "https://github.com/hoxnox/snappystream/archive/refs/tags/1.0.0.tar.gz"
	snappystream_cmd_tar := exec.Command("curl", "-L", snappystream_tar_url, "-o", "package.tar.gz")
	err := snappystream_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snappystream_zip_url := "https://github.com/hoxnox/snappystream/archive/refs/tags/1.0.0.zip"
	snappystream_cmd_zip := exec.Command("curl", "-L", snappystream_zip_url, "-o", "package.zip")
	err = snappystream_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snappystream_bin_url := "https://github.com/hoxnox/snappystream/archive/refs/tags/1.0.0.bin"
	snappystream_cmd_bin := exec.Command("curl", "-L", snappystream_bin_url, "-o", "binary.bin")
	err = snappystream_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snappystream_src_url := "https://github.com/hoxnox/snappystream/archive/refs/tags/1.0.0.src.tar.gz"
	snappystream_cmd_src := exec.Command("curl", "-L", snappystream_src_url, "-o", "source.tar.gz")
	err = snappystream_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snappystream_cmd_direct := exec.Command("./binary")
	err = snappystream_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
}
