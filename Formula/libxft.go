package main

// Libxft - X.Org: X FreeType library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxft() {
	// Método 1: Descargar y extraer .tar.gz
	libxft_tar_url := "https://www.x.org/archive/individual/lib/libXft-2.3.8.tar.xz"
	libxft_cmd_tar := exec.Command("curl", "-L", libxft_tar_url, "-o", "package.tar.gz")
	err := libxft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxft_zip_url := "https://www.x.org/archive/individual/lib/libXft-2.3.8.tar.xz"
	libxft_cmd_zip := exec.Command("curl", "-L", libxft_zip_url, "-o", "package.zip")
	err = libxft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxft_bin_url := "https://www.x.org/archive/individual/lib/libXft-2.3.8.tar.xz"
	libxft_cmd_bin := exec.Command("curl", "-L", libxft_bin_url, "-o", "binary.bin")
	err = libxft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxft_src_url := "https://www.x.org/archive/individual/lib/libXft-2.3.8.tar.xz"
	libxft_cmd_src := exec.Command("curl", "-L", libxft_src_url, "-o", "source.tar.gz")
	err = libxft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxft_cmd_direct := exec.Command("./binary")
	err = libxft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
}
