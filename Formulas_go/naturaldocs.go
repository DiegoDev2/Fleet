package main

// Naturaldocs - Extensible, multi-language documentation generator
// Homepage: https://www.naturaldocs.org/

import (
	"fmt"
	
	"os/exec"
)

func installNaturaldocs() {
	// Método 1: Descargar y extraer .tar.gz
	naturaldocs_tar_url := "https://downloads.sourceforge.net/project/naturaldocs/Stable%20Releases/2.3/Natural_Docs_2.3.zip"
	naturaldocs_cmd_tar := exec.Command("curl", "-L", naturaldocs_tar_url, "-o", "package.tar.gz")
	err := naturaldocs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	naturaldocs_zip_url := "https://downloads.sourceforge.net/project/naturaldocs/Stable%20Releases/2.3/Natural_Docs_2.3.zip"
	naturaldocs_cmd_zip := exec.Command("curl", "-L", naturaldocs_zip_url, "-o", "package.zip")
	err = naturaldocs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	naturaldocs_bin_url := "https://downloads.sourceforge.net/project/naturaldocs/Stable%20Releases/2.3/Natural_Docs_2.3.zip"
	naturaldocs_cmd_bin := exec.Command("curl", "-L", naturaldocs_bin_url, "-o", "binary.bin")
	err = naturaldocs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	naturaldocs_src_url := "https://downloads.sourceforge.net/project/naturaldocs/Stable%20Releases/2.3/Natural_Docs_2.3.zip"
	naturaldocs_cmd_src := exec.Command("curl", "-L", naturaldocs_src_url, "-o", "source.tar.gz")
	err = naturaldocs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	naturaldocs_cmd_direct := exec.Command("./binary")
	err = naturaldocs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mono")
exec.Command("latte", "install", "mono")
}
