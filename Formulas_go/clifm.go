package main

// Clifm - Command-line Interface File Manager
// Homepage: https://github.com/leo-arch/clifm

import (
	"fmt"
	
	"os/exec"
)

func installClifm() {
	// Método 1: Descargar y extraer .tar.gz
	clifm_tar_url := "https://github.com/leo-arch/clifm/archive/refs/tags/v1.20.tar.gz"
	clifm_cmd_tar := exec.Command("curl", "-L", clifm_tar_url, "-o", "package.tar.gz")
	err := clifm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clifm_zip_url := "https://github.com/leo-arch/clifm/archive/refs/tags/v1.20.zip"
	clifm_cmd_zip := exec.Command("curl", "-L", clifm_zip_url, "-o", "package.zip")
	err = clifm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clifm_bin_url := "https://github.com/leo-arch/clifm/archive/refs/tags/v1.20.bin"
	clifm_cmd_bin := exec.Command("curl", "-L", clifm_bin_url, "-o", "binary.bin")
	err = clifm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clifm_src_url := "https://github.com/leo-arch/clifm/archive/refs/tags/v1.20.src.tar.gz"
	clifm_cmd_src := exec.Command("curl", "-L", clifm_src_url, "-o", "source.tar.gz")
	err = clifm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clifm_cmd_direct := exec.Command("./binary")
	err = clifm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: acl")
exec.Command("latte", "install", "acl")
	fmt.Println("Instalando dependencia: libcap")
exec.Command("latte", "install", "libcap")
}
