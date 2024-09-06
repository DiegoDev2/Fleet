package main

// Solid - Collision detection library for geometric objects in 3D space
// Homepage: https://github.com/dtecta/solid3/

import (
	"fmt"
	
	"os/exec"
)

func installSolid() {
	// Método 1: Descargar y extraer .tar.gz
	solid_tar_url := "https://github.com/dtecta/solid3/archive/ec3e218616749949487f81165f8b478b16bc7932.tar.gz"
	solid_cmd_tar := exec.Command("curl", "-L", solid_tar_url, "-o", "package.tar.gz")
	err := solid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solid_zip_url := "https://github.com/dtecta/solid3/archive/ec3e218616749949487f81165f8b478b16bc7932.zip"
	solid_cmd_zip := exec.Command("curl", "-L", solid_zip_url, "-o", "package.zip")
	err = solid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solid_bin_url := "https://github.com/dtecta/solid3/archive/ec3e218616749949487f81165f8b478b16bc7932.bin"
	solid_cmd_bin := exec.Command("curl", "-L", solid_bin_url, "-o", "binary.bin")
	err = solid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solid_src_url := "https://github.com/dtecta/solid3/archive/ec3e218616749949487f81165f8b478b16bc7932.src.tar.gz"
	solid_cmd_src := exec.Command("curl", "-L", solid_src_url, "-o", "source.tar.gz")
	err = solid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solid_cmd_direct := exec.Command("./binary")
	err = solid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
