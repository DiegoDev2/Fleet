package main

// Uwsgi - Full stack for building hosting services
// Homepage: https://uwsgi-docs.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installUwsgi() {
	// Método 1: Descargar y extraer .tar.gz
	uwsgi_tar_url := "https://files.pythonhosted.org/packages/3a/7a/4c910bdc9d32640ba89f8d1dc256872c2b5e64830759f7dc346815f5b3b1/uwsgi-2.0.26.tar.gz"
	uwsgi_cmd_tar := exec.Command("curl", "-L", uwsgi_tar_url, "-o", "package.tar.gz")
	err := uwsgi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uwsgi_zip_url := "https://files.pythonhosted.org/packages/3a/7a/4c910bdc9d32640ba89f8d1dc256872c2b5e64830759f7dc346815f5b3b1/uwsgi-2.0.26.zip"
	uwsgi_cmd_zip := exec.Command("curl", "-L", uwsgi_zip_url, "-o", "package.zip")
	err = uwsgi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uwsgi_bin_url := "https://files.pythonhosted.org/packages/3a/7a/4c910bdc9d32640ba89f8d1dc256872c2b5e64830759f7dc346815f5b3b1/uwsgi-2.0.26.bin"
	uwsgi_cmd_bin := exec.Command("curl", "-L", uwsgi_bin_url, "-o", "binary.bin")
	err = uwsgi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uwsgi_src_url := "https://files.pythonhosted.org/packages/3a/7a/4c910bdc9d32640ba89f8d1dc256872c2b5e64830759f7dc346815f5b3b1/uwsgi-2.0.26.src.tar.gz"
	uwsgi_cmd_src := exec.Command("curl", "-L", uwsgi_src_url, "-o", "source.tar.gz")
	err = uwsgi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uwsgi_cmd_direct := exec.Command("./binary")
	err = uwsgi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: yajl")
exec.Command("latte", "install", "yajl")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
