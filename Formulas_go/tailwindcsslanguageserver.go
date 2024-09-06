package main

// TailwindcssLanguageServer - LSP for TailwindCSS
// Homepage: https://github.com/tailwindlabs/tailwindcss-intellisense/tree/HEAD/packages/tailwindcss-language-server

import (
	"fmt"
	
	"os/exec"
)

func installTailwindcssLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	tailwindcsslanguageserver_tar_url := "https://registry.npmjs.org/@tailwindcss/language-server/-/language-server-0.0.24.tgz"
	tailwindcsslanguageserver_cmd_tar := exec.Command("curl", "-L", tailwindcsslanguageserver_tar_url, "-o", "package.tar.gz")
	err := tailwindcsslanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tailwindcsslanguageserver_zip_url := "https://registry.npmjs.org/@tailwindcss/language-server/-/language-server-0.0.24.tgz"
	tailwindcsslanguageserver_cmd_zip := exec.Command("curl", "-L", tailwindcsslanguageserver_zip_url, "-o", "package.zip")
	err = tailwindcsslanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tailwindcsslanguageserver_bin_url := "https://registry.npmjs.org/@tailwindcss/language-server/-/language-server-0.0.24.tgz"
	tailwindcsslanguageserver_cmd_bin := exec.Command("curl", "-L", tailwindcsslanguageserver_bin_url, "-o", "binary.bin")
	err = tailwindcsslanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tailwindcsslanguageserver_src_url := "https://registry.npmjs.org/@tailwindcss/language-server/-/language-server-0.0.24.tgz"
	tailwindcsslanguageserver_cmd_src := exec.Command("curl", "-L", tailwindcsslanguageserver_src_url, "-o", "source.tar.gz")
	err = tailwindcsslanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tailwindcsslanguageserver_cmd_direct := exec.Command("./binary")
	err = tailwindcsslanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
