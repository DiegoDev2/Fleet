package main

// Sispmctl - Control Gembird SIS-PM programmable power outlet strips
// Homepage: https://sispmctl.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSispmctl() {
	// Método 1: Descargar y extraer .tar.gz
	sispmctl_tar_url := "https://downloads.sourceforge.net/project/sispmctl/sispmctl/sispmctl-4.12/sispmctl-4.12.tar.gz"
	sispmctl_cmd_tar := exec.Command("curl", "-L", sispmctl_tar_url, "-o", "package.tar.gz")
	err := sispmctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sispmctl_zip_url := "https://downloads.sourceforge.net/project/sispmctl/sispmctl/sispmctl-4.12/sispmctl-4.12.zip"
	sispmctl_cmd_zip := exec.Command("curl", "-L", sispmctl_zip_url, "-o", "package.zip")
	err = sispmctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sispmctl_bin_url := "https://downloads.sourceforge.net/project/sispmctl/sispmctl/sispmctl-4.12/sispmctl-4.12.bin"
	sispmctl_cmd_bin := exec.Command("curl", "-L", sispmctl_bin_url, "-o", "binary.bin")
	err = sispmctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sispmctl_src_url := "https://downloads.sourceforge.net/project/sispmctl/sispmctl/sispmctl-4.12/sispmctl-4.12.src.tar.gz"
	sispmctl_cmd_src := exec.Command("curl", "-L", sispmctl_src_url, "-o", "source.tar.gz")
	err = sispmctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sispmctl_cmd_direct := exec.Command("./binary")
	err = sispmctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb-compat")
	exec.Command("latte", "install", "libusb-compat").Run()
}
