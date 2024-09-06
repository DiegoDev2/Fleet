package main

// Libice - X.Org: Inter-Client Exchange Library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibice() {
	// Método 1: Descargar y extraer .tar.gz
	libice_tar_url := "https://www.x.org/archive/individual/lib/libICE-1.1.1.tar.xz"
	libice_cmd_tar := exec.Command("curl", "-L", libice_tar_url, "-o", "package.tar.gz")
	err := libice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libice_zip_url := "https://www.x.org/archive/individual/lib/libICE-1.1.1.tar.xz"
	libice_cmd_zip := exec.Command("curl", "-L", libice_zip_url, "-o", "package.zip")
	err = libice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libice_bin_url := "https://www.x.org/archive/individual/lib/libICE-1.1.1.tar.xz"
	libice_cmd_bin := exec.Command("curl", "-L", libice_bin_url, "-o", "binary.bin")
	err = libice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libice_src_url := "https://www.x.org/archive/individual/lib/libICE-1.1.1.tar.xz"
	libice_cmd_src := exec.Command("curl", "-L", libice_src_url, "-o", "source.tar.gz")
	err = libice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libice_cmd_direct := exec.Command("./binary")
	err = libice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xtrans")
exec.Command("latte", "install", "xtrans")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
