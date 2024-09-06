package main

// NetlifyCli - Netlify command-line tool
// Homepage: https://www.netlify.com/docs/cli

import (
	"fmt"
	
	"os/exec"
)

func installNetlifyCli() {
	// Método 1: Descargar y extraer .tar.gz
	netlifycli_tar_url := "https://registry.npmjs.org/netlify-cli/-/netlify-cli-17.34.3.tgz"
	netlifycli_cmd_tar := exec.Command("curl", "-L", netlifycli_tar_url, "-o", "package.tar.gz")
	err := netlifycli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netlifycli_zip_url := "https://registry.npmjs.org/netlify-cli/-/netlify-cli-17.34.3.tgz"
	netlifycli_cmd_zip := exec.Command("curl", "-L", netlifycli_zip_url, "-o", "package.zip")
	err = netlifycli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netlifycli_bin_url := "https://registry.npmjs.org/netlify-cli/-/netlify-cli-17.34.3.tgz"
	netlifycli_cmd_bin := exec.Command("curl", "-L", netlifycli_bin_url, "-o", "binary.bin")
	err = netlifycli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netlifycli_src_url := "https://registry.npmjs.org/netlify-cli/-/netlify-cli-17.34.3.tgz"
	netlifycli_cmd_src := exec.Command("curl", "-L", netlifycli_src_url, "-o", "source.tar.gz")
	err = netlifycli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netlifycli_cmd_direct := exec.Command("./binary")
	err = netlifycli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: vips")
exec.Command("latte", "install", "vips")
	fmt.Println("Instalando dependencia: xsel")
exec.Command("latte", "install", "xsel")
}
