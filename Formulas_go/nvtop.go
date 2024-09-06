package main

// Nvtop - Interactive GPU process monitor
// Homepage: https://github.com/Syllo/nvtop

import (
	"fmt"
	
	"os/exec"
)

func installNvtop() {
	// Método 1: Descargar y extraer .tar.gz
	nvtop_tar_url := "https://github.com/Syllo/nvtop/archive/refs/tags/3.1.0.tar.gz"
	nvtop_cmd_tar := exec.Command("curl", "-L", nvtop_tar_url, "-o", "package.tar.gz")
	err := nvtop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nvtop_zip_url := "https://github.com/Syllo/nvtop/archive/refs/tags/3.1.0.zip"
	nvtop_cmd_zip := exec.Command("curl", "-L", nvtop_zip_url, "-o", "package.zip")
	err = nvtop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nvtop_bin_url := "https://github.com/Syllo/nvtop/archive/refs/tags/3.1.0.bin"
	nvtop_cmd_bin := exec.Command("curl", "-L", nvtop_bin_url, "-o", "binary.bin")
	err = nvtop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nvtop_src_url := "https://github.com/Syllo/nvtop/archive/refs/tags/3.1.0.src.tar.gz"
	nvtop_cmd_src := exec.Command("curl", "-L", nvtop_src_url, "-o", "source.tar.gz")
	err = nvtop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nvtop_cmd_direct := exec.Command("./binary")
	err = nvtop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libdrm")
exec.Command("latte", "install", "libdrm")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
