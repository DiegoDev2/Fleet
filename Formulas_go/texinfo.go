package main

// Texinfo - Official documentation format of the GNU project
// Homepage: https://www.gnu.org/software/texinfo/

import (
	"fmt"
	
	"os/exec"
)

func installTexinfo() {
	// Método 1: Descargar y extraer .tar.gz
	texinfo_tar_url := "https://ftp.gnu.org/gnu/texinfo/texinfo-7.1.tar.xz"
	texinfo_cmd_tar := exec.Command("curl", "-L", texinfo_tar_url, "-o", "package.tar.gz")
	err := texinfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texinfo_zip_url := "https://ftp.gnu.org/gnu/texinfo/texinfo-7.1.tar.xz"
	texinfo_cmd_zip := exec.Command("curl", "-L", texinfo_zip_url, "-o", "package.zip")
	err = texinfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texinfo_bin_url := "https://ftp.gnu.org/gnu/texinfo/texinfo-7.1.tar.xz"
	texinfo_cmd_bin := exec.Command("curl", "-L", texinfo_bin_url, "-o", "binary.bin")
	err = texinfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texinfo_src_url := "https://ftp.gnu.org/gnu/texinfo/texinfo-7.1.tar.xz"
	texinfo_cmd_src := exec.Command("curl", "-L", texinfo_src_url, "-o", "source.tar.gz")
	err = texinfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texinfo_cmd_direct := exec.Command("./binary")
	err = texinfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
