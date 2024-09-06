package main

// Werf - Consistent delivery tool for Kubernetes
// Homepage: https://werf.io/

import (
	"fmt"
	
	"os/exec"
)

func installWerf() {
	// Método 1: Descargar y extraer .tar.gz
	werf_tar_url := "https://github.com/werf/werf/archive/refs/tags/v2.10.5.tar.gz"
	werf_cmd_tar := exec.Command("curl", "-L", werf_tar_url, "-o", "package.tar.gz")
	err := werf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	werf_zip_url := "https://github.com/werf/werf/archive/refs/tags/v2.10.5.zip"
	werf_cmd_zip := exec.Command("curl", "-L", werf_zip_url, "-o", "package.zip")
	err = werf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	werf_bin_url := "https://github.com/werf/werf/archive/refs/tags/v2.10.5.bin"
	werf_cmd_bin := exec.Command("curl", "-L", werf_bin_url, "-o", "binary.bin")
	err = werf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	werf_src_url := "https://github.com/werf/werf/archive/refs/tags/v2.10.5.src.tar.gz"
	werf_cmd_src := exec.Command("curl", "-L", werf_src_url, "-o", "source.tar.gz")
	err = werf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	werf_cmd_direct := exec.Command("./binary")
	err = werf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: btrfs-progs")
exec.Command("latte", "install", "btrfs-progs")
	fmt.Println("Instalando dependencia: device-mapper")
exec.Command("latte", "install", "device-mapper")
}
