package main

// Dwarfutils - Dump and produce DWARF debug information in ELF objects
// Homepage: https://www.prevanders.net/dwarf.html

import (
	"fmt"
	
	"os/exec"
)

func installDwarfutils() {
	// Método 1: Descargar y extraer .tar.gz
	dwarfutils_tar_url := "https://www.prevanders.net/libdwarf-0.11.0.tar.xz"
	dwarfutils_cmd_tar := exec.Command("curl", "-L", dwarfutils_tar_url, "-o", "package.tar.gz")
	err := dwarfutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dwarfutils_zip_url := "https://www.prevanders.net/libdwarf-0.11.0.tar.xz"
	dwarfutils_cmd_zip := exec.Command("curl", "-L", dwarfutils_zip_url, "-o", "package.zip")
	err = dwarfutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dwarfutils_bin_url := "https://www.prevanders.net/libdwarf-0.11.0.tar.xz"
	dwarfutils_cmd_bin := exec.Command("curl", "-L", dwarfutils_bin_url, "-o", "binary.bin")
	err = dwarfutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dwarfutils_src_url := "https://www.prevanders.net/libdwarf-0.11.0.tar.xz"
	dwarfutils_cmd_src := exec.Command("curl", "-L", dwarfutils_src_url, "-o", "source.tar.gz")
	err = dwarfutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dwarfutils_cmd_direct := exec.Command("./binary")
	err = dwarfutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
