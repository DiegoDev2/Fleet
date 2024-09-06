package main

// Btop - Resource monitor. C++ version and continuation of bashtop and bpytop
// Homepage: https://github.com/aristocratos/btop

import (
	"fmt"
	
	"os/exec"
)

func installBtop() {
	// Método 1: Descargar y extraer .tar.gz
	btop_tar_url := "https://github.com/aristocratos/btop/archive/refs/tags/v1.3.2.tar.gz"
	btop_cmd_tar := exec.Command("curl", "-L", btop_tar_url, "-o", "package.tar.gz")
	err := btop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	btop_zip_url := "https://github.com/aristocratos/btop/archive/refs/tags/v1.3.2.zip"
	btop_cmd_zip := exec.Command("curl", "-L", btop_zip_url, "-o", "package.zip")
	err = btop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	btop_bin_url := "https://github.com/aristocratos/btop/archive/refs/tags/v1.3.2.bin"
	btop_cmd_bin := exec.Command("curl", "-L", btop_bin_url, "-o", "binary.bin")
	err = btop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	btop_src_url := "https://github.com/aristocratos/btop/archive/refs/tags/v1.3.2.src.tar.gz"
	btop_cmd_src := exec.Command("curl", "-L", btop_src_url, "-o", "source.tar.gz")
	err = btop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	btop_cmd_direct := exec.Command("./binary")
	err = btop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
