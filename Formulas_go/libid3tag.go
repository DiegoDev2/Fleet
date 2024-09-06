package main

// Libid3tag - ID3 tag manipulation library
// Homepage: https://www.underbit.com/products/mad/

import (
	"fmt"
	
	"os/exec"
)

func installLibid3tag() {
	// Método 1: Descargar y extraer .tar.gz
	libid3tag_tar_url := "https://codeberg.org/tenacityteam/libid3tag/archive/0.16.3.tar.gz"
	libid3tag_cmd_tar := exec.Command("curl", "-L", libid3tag_tar_url, "-o", "package.tar.gz")
	err := libid3tag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libid3tag_zip_url := "https://codeberg.org/tenacityteam/libid3tag/archive/0.16.3.zip"
	libid3tag_cmd_zip := exec.Command("curl", "-L", libid3tag_zip_url, "-o", "package.zip")
	err = libid3tag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libid3tag_bin_url := "https://codeberg.org/tenacityteam/libid3tag/archive/0.16.3.bin"
	libid3tag_cmd_bin := exec.Command("curl", "-L", libid3tag_bin_url, "-o", "binary.bin")
	err = libid3tag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libid3tag_src_url := "https://codeberg.org/tenacityteam/libid3tag/archive/0.16.3.src.tar.gz"
	libid3tag_cmd_src := exec.Command("curl", "-L", libid3tag_src_url, "-o", "source.tar.gz")
	err = libid3tag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libid3tag_cmd_direct := exec.Command("./binary")
	err = libid3tag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
