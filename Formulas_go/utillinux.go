package main

// UtilLinux - Collection of Linux utilities
// Homepage: https://github.com/util-linux/util-linux

import (
	"fmt"
	
	"os/exec"
)

func installUtilLinux() {
	// Método 1: Descargar y extraer .tar.gz
	utillinux_tar_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	utillinux_cmd_tar := exec.Command("curl", "-L", utillinux_tar_url, "-o", "package.tar.gz")
	err := utillinux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	utillinux_zip_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	utillinux_cmd_zip := exec.Command("curl", "-L", utillinux_zip_url, "-o", "package.zip")
	err = utillinux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	utillinux_bin_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	utillinux_cmd_bin := exec.Command("curl", "-L", utillinux_bin_url, "-o", "binary.bin")
	err = utillinux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	utillinux_src_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	utillinux_cmd_src := exec.Command("curl", "-L", utillinux_src_url, "-o", "source.tar.gz")
	err = utillinux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	utillinux_cmd_direct := exec.Command("./binary")
	err = utillinux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
