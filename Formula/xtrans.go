package main

// Xtrans - X.Org: X Network Transport layer shared code
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installXtrans() {
	// Método 1: Descargar y extraer .tar.gz
	xtrans_tar_url := "https://www.x.org/archive/individual/lib/xtrans-1.5.0.tar.xz"
	xtrans_cmd_tar := exec.Command("curl", "-L", xtrans_tar_url, "-o", "package.tar.gz")
	err := xtrans_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xtrans_zip_url := "https://www.x.org/archive/individual/lib/xtrans-1.5.0.tar.xz"
	xtrans_cmd_zip := exec.Command("curl", "-L", xtrans_zip_url, "-o", "package.zip")
	err = xtrans_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xtrans_bin_url := "https://www.x.org/archive/individual/lib/xtrans-1.5.0.tar.xz"
	xtrans_cmd_bin := exec.Command("curl", "-L", xtrans_bin_url, "-o", "binary.bin")
	err = xtrans_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xtrans_src_url := "https://www.x.org/archive/individual/lib/xtrans-1.5.0.tar.xz"
	xtrans_cmd_src := exec.Command("curl", "-L", xtrans_src_url, "-o", "source.tar.gz")
	err = xtrans_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xtrans_cmd_direct := exec.Command("./binary")
	err = xtrans_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
