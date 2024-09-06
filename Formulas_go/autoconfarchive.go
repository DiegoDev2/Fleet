package main

// AutoconfArchive - Collection of over 500 reusable autoconf macros
// Homepage: https://savannah.gnu.org/projects/autoconf-archive/

import (
	"fmt"
	
	"os/exec"
)

func installAutoconfArchive() {
	// Método 1: Descargar y extraer .tar.gz
	autoconfarchive_tar_url := "https://ftp.gnu.org/gnu/autoconf-archive/autoconf-archive-2023.02.20.tar.xz"
	autoconfarchive_cmd_tar := exec.Command("curl", "-L", autoconfarchive_tar_url, "-o", "package.tar.gz")
	err := autoconfarchive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autoconfarchive_zip_url := "https://ftp.gnu.org/gnu/autoconf-archive/autoconf-archive-2023.02.20.tar.xz"
	autoconfarchive_cmd_zip := exec.Command("curl", "-L", autoconfarchive_zip_url, "-o", "package.zip")
	err = autoconfarchive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autoconfarchive_bin_url := "https://ftp.gnu.org/gnu/autoconf-archive/autoconf-archive-2023.02.20.tar.xz"
	autoconfarchive_cmd_bin := exec.Command("curl", "-L", autoconfarchive_bin_url, "-o", "binary.bin")
	err = autoconfarchive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autoconfarchive_src_url := "https://ftp.gnu.org/gnu/autoconf-archive/autoconf-archive-2023.02.20.tar.xz"
	autoconfarchive_cmd_src := exec.Command("curl", "-L", autoconfarchive_src_url, "-o", "source.tar.gz")
	err = autoconfarchive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autoconfarchive_cmd_direct := exec.Command("./binary")
	err = autoconfarchive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}
