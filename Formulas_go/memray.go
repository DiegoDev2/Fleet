package main

// Memray - Memory profiler for Python applications
// Homepage: https://bloomberg.github.io/memray/

import (
	"fmt"
	
	"os/exec"
)

func installMemray() {
	// Método 1: Descargar y extraer .tar.gz
	memray_tar_url := "https://files.pythonhosted.org/packages/88/8b/0a9854e5b6ce0875f2e2ad163cfecdb8de59fc1a2f1b9a1cb7c683a67826/memray-1.12.0.tar.gz"
	memray_cmd_tar := exec.Command("curl", "-L", memray_tar_url, "-o", "package.tar.gz")
	err := memray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	memray_zip_url := "https://files.pythonhosted.org/packages/88/8b/0a9854e5b6ce0875f2e2ad163cfecdb8de59fc1a2f1b9a1cb7c683a67826/memray-1.12.0.zip"
	memray_cmd_zip := exec.Command("curl", "-L", memray_zip_url, "-o", "package.zip")
	err = memray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	memray_bin_url := "https://files.pythonhosted.org/packages/88/8b/0a9854e5b6ce0875f2e2ad163cfecdb8de59fc1a2f1b9a1cb7c683a67826/memray-1.12.0.bin"
	memray_cmd_bin := exec.Command("curl", "-L", memray_bin_url, "-o", "binary.bin")
	err = memray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	memray_src_url := "https://files.pythonhosted.org/packages/88/8b/0a9854e5b6ce0875f2e2ad163cfecdb8de59fc1a2f1b9a1cb7c683a67826/memray-1.12.0.src.tar.gz"
	memray_cmd_src := exec.Command("curl", "-L", memray_src_url, "-o", "source.tar.gz")
	err = memray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	memray_cmd_direct := exec.Command("./binary")
	err = memray_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: libunwind")
exec.Command("latte", "install", "libunwind")
}
