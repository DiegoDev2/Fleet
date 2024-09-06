package main

// Kmod - Linux kernel module handling
// Homepage: https://git.kernel.org/pub/scm/utils/kernel/kmod/kmod.git

import (
	"fmt"
	
	"os/exec"
)

func installKmod() {
	// Método 1: Descargar y extraer .tar.gz
	kmod_tar_url := "https://mirrors.edge.kernel.org/pub/linux/utils/kernel/kmod/kmod-33.tar.xz"
	kmod_cmd_tar := exec.Command("curl", "-L", kmod_tar_url, "-o", "package.tar.gz")
	err := kmod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kmod_zip_url := "https://mirrors.edge.kernel.org/pub/linux/utils/kernel/kmod/kmod-33.tar.xz"
	kmod_cmd_zip := exec.Command("curl", "-L", kmod_zip_url, "-o", "package.zip")
	err = kmod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kmod_bin_url := "https://mirrors.edge.kernel.org/pub/linux/utils/kernel/kmod/kmod-33.tar.xz"
	kmod_cmd_bin := exec.Command("curl", "-L", kmod_bin_url, "-o", "binary.bin")
	err = kmod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kmod_src_url := "https://mirrors.edge.kernel.org/pub/linux/utils/kernel/kmod/kmod-33.tar.xz"
	kmod_cmd_src := exec.Command("curl", "-L", kmod_src_url, "-o", "source.tar.gz")
	err = kmod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kmod_cmd_direct := exec.Command("./binary")
	err = kmod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
}
