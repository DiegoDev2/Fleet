package main

// Wiredtiger - High performance NoSQL extensible platform for data management
// Homepage: https://source.wiredtiger.com/

import (
	"fmt"
	
	"os/exec"
)

func installWiredtiger() {
	// Método 1: Descargar y extraer .tar.gz
	wiredtiger_tar_url := "https://github.com/wiredtiger/wiredtiger/archive/refs/tags/11.2.0.tar.gz"
	wiredtiger_cmd_tar := exec.Command("curl", "-L", wiredtiger_tar_url, "-o", "package.tar.gz")
	err := wiredtiger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wiredtiger_zip_url := "https://github.com/wiredtiger/wiredtiger/archive/refs/tags/11.2.0.zip"
	wiredtiger_cmd_zip := exec.Command("curl", "-L", wiredtiger_zip_url, "-o", "package.zip")
	err = wiredtiger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wiredtiger_bin_url := "https://github.com/wiredtiger/wiredtiger/archive/refs/tags/11.2.0.bin"
	wiredtiger_cmd_bin := exec.Command("curl", "-L", wiredtiger_bin_url, "-o", "binary.bin")
	err = wiredtiger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wiredtiger_src_url := "https://github.com/wiredtiger/wiredtiger/archive/refs/tags/11.2.0.src.tar.gz"
	wiredtiger_cmd_src := exec.Command("curl", "-L", wiredtiger_src_url, "-o", "source.tar.gz")
	err = wiredtiger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wiredtiger_cmd_direct := exec.Command("./binary")
	err = wiredtiger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ccache")
exec.Command("latte", "install", "ccache")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
