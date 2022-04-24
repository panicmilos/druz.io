from jinja2 import Environment, PackageLoader, select_autoescape


class TemplateManager:
  def __init__(self, package_name = 'main'):
    self._env = Environment(
      loader=PackageLoader(package_name),
      autoescape=select_autoescape()
    )

  def get_template(self, template_name):
    return self._env.get_template(template_name + '.html')

  def compute_tempalte(self, template_name, params = {}):
    template = self.get_template(template_name)

    return template.render(params)
